package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
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

func (this *RBACCtl) Build(athena *athena.Athena) {
	// role列表
	athena.Handle("GET", "/roles", this.roles)
}
