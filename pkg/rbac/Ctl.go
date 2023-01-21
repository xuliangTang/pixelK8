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
	RoleSvc           *RoleService           `inject:"-"`
	RoleBindingSvc    *RoleBindingService    `inject:"-"`
	ServiceAccountSvc *ServiceAccountService `inject:"-"`
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

func (this *RBACCtl) rolesAll(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	roleList := this.RoleSvc.ListByNs(ns)

	return roleList
}

func (this *RBACCtl) createRole(ctx *gin.Context) any {
	role := &rbacV1.Role{}
	athena.Error(ctx.BindJSON(role))
	athena.Error(this.RoleSvc.Create(role))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return role
}

func (this *RBACCtl) showRole(ctx *gin.Context) any {
	uri := &struct {
		Namespace string `uri:"ns" binding:"required"`
		RoleName  string `uri:"role" binding:"required"`
	}{}

	athena.Error(ctx.BindUri(uri))

	return this.RoleSvc.Show(uri.Namespace, uri.RoleName)
}

func (this *RBACCtl) updateRole(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		Namespace string `uri:"ns" binding:"required"`
		RoleName  string `uri:"role" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))

	role := &rbacV1.Role{}
	athena.Error(ctx.BindJSON(role))

	athena.Error(this.RoleSvc.Update(uri.Namespace, uri.RoleName, role.Rules))

	return
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

func (this *RBACCtl) roleBindings(ctx *gin.Context) athena.Collection {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	page := athena.NewPageWithCtx(ctx)
	roleBindingList := this.RoleBindingSvc.ListByNs(ns)

	return this.RoleBindingSvc.Paging(page, roleBindingList)
}

func (this *RBACCtl) createRoleBinding(ctx *gin.Context) any {
	roleBinding := &rbacV1.RoleBinding{}
	athena.Error(ctx.BindJSON(roleBinding))
	athena.Error(this.RoleBindingSvc.Create(roleBinding))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return roleBinding
}

func (this *RBACCtl) deleteRoleBinding(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		Namespace       string `uri:"ns" binding:"required"`
		RoleBindingName string `uri:"roleBinding" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.RoleBindingSvc.Delete(uri.Namespace, uri.RoleBindingName))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) addUserToRoleBinding(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		Namespace       string `uri:"ns" binding:"required"`
		RoleBindingName string `uri:"roleBinding" binding:"required"`
	}{}
	subject := &rbacV1.Subject{}

	athena.Error(ctx.BindUri(uri))
	athena.Error(ctx.BindJSON(subject))
	athena.Error(this.RoleBindingSvc.AddUser(uri.Namespace, uri.RoleBindingName, subject))

	return
}

func (this *RBACCtl) removeUserFromRoleBinding(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		Namespace       string `uri:"ns" binding:"required"`
		RoleBindingName string `uri:"roleBinding" binding:"required"`
	}{}
	subject := &rbacV1.Subject{}

	athena.Error(ctx.BindUri(uri))
	athena.Error(ctx.BindJSON(subject))
	athena.Error(this.RoleBindingSvc.RemoveUser(uri.Namespace, uri.RoleBindingName, subject))

	return
}

func (this *RBACCtl) serviceAccounts(ctx *gin.Context) athena.Collection {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	page := athena.NewPageWithCtx(ctx)
	saList := this.ServiceAccountSvc.ListByNs(ns)

	return this.ServiceAccountSvc.Paging(page, saList)
}

func (this *RBACCtl) Build(athena *athena.Athena) {
	// role列表
	athena.Handle("GET", "/roles", this.roles)
	// 获取全部role
	athena.Handle("GET", "/roles/all", this.rolesAll)
	// 创建role
	athena.Handle("POST", "/role", this.createRole)
	// 查看role
	athena.Handle("GET", "/role/:ns/:role", this.showRole)
	// 编辑role
	athena.Handle("PUT", "/role/:ns/:role", this.updateRole)
	// 删除role
	athena.Handle("DELETE", "/role/:ns/:role", this.deleteRole)
	// roleBinding列表
	athena.Handle("GET", "/roleBindings", this.roleBindings)
	// 创建roleBinding
	athena.Handle("POST", "/roleBinding", this.createRoleBinding)
	// 删除roleBinding
	athena.Handle("DELETE", "/roleBinding/:ns/:roleBinding", this.deleteRoleBinding)
	// roleBinding增加user
	athena.Handle("PATCH", "/roleBinding/:ns/:roleBinding/user", this.addUserToRoleBinding)
	// roleBinding移除user
	athena.Handle("PATCH", "/roleBinding/:ns/:roleBinding/user/remove", this.removeUserFromRoleBinding)
	// serviceAccount列表
	athena.Handle("GET", "/serviceAccounts", this.serviceAccounts)
}
