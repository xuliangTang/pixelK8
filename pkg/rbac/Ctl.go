package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	"net/http"
	"pixelk8/src/properties"
)

// RBACCtl @Controller
type RBACCtl struct {
	RoleSvc *RoleService `inject:"-"`
}

func NewRBACCtl() *RBACCtl {
	return &RBACCtl{}
}

func (this *RBACCtl) roles(ctx *gin.Context) athena.Collection {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	page := athena.NewPageWithCtx(ctx)
	roleList := this.RoleSvc.ListByNs(ns)

	return this.RoleSvc.Paging(page, roleList)
}

func (this *RBACCtl) createRole(ctx *gin.Context) any {
	role := &rbacV1.Role{}
	athena.Error(ctx.BindJSON(role))
	athena.Error(this.RoleSvc.Create(role))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return role
}

func (this *RBACCtl) deleteRole(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		Namespace string `uri:"ns" binding:"required"`
		RoleName  string `uri:"role" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.RoleSvc.Delete(uri.Namespace, uri.RoleName))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) Build(athena *athena.Athena) {
	// role列表
	athena.Handle("GET", "/roles", this.roles)
	// 创建role
	athena.Handle("POST", "/role", this.createRole)
	// 删除role
	athena.Handle("DELETE", "/role/:ns/:role", this.deleteRole)
}
