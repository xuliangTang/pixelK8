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
	RoleSvc               *RoleService               `inject:"-"`
	RoleBindingSvc        *RoleBindingService        `inject:"-"`
	ServiceAccountSvc     *ServiceAccountService     `inject:"-"`
	ClusterRoleSvc        *ClusterRoleService        `inject:"-"`
	ClusterRoleBindingSvc *ClusterRoleBindingService `inject:"-"`
	UserAccountService    *UserAccountService        `inject:"-"`
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

func (this *RBACCtl) deleteServiceAccount(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		Namespace          string `uri:"ns" binding:"required"`
		ServiceAccountName string `uri:"serviceAccount" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.ServiceAccountSvc.Delete(uri.Namespace, uri.ServiceAccountName))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) clusterRoles(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	clusterRoleList := this.ClusterRoleSvc.List()

	return this.ClusterRoleSvc.Paging(page, clusterRoleList)
}

func (this *RBACCtl) clusterRolesAll(ctx *gin.Context) any {
	clusterRolesList := this.ClusterRoleSvc.List()

	return clusterRolesList
}

func (this *RBACCtl) showClusterRole(ctx *gin.Context) any {
	uri := &struct {
		ClusterRoleName string `uri:"clusterRole" binding:"required"`
	}{}

	athena.Error(ctx.BindUri(uri))

	return this.ClusterRoleSvc.Show(uri.ClusterRoleName)
}

func (this *RBACCtl) createClusterRole(ctx *gin.Context) any {
	clusterRole := &rbacV1.ClusterRole{}
	athena.Error(ctx.BindJSON(clusterRole))
	athena.Error(this.ClusterRoleSvc.Create(clusterRole))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return clusterRole
}

func (this *RBACCtl) updateClusterRole(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		ClusterRoleName string `uri:"clusterRole" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))

	clusterRole := &rbacV1.ClusterRole{}
	athena.Error(ctx.BindJSON(clusterRole))

	athena.Error(this.ClusterRoleSvc.Update(uri.ClusterRoleName, clusterRole.Rules))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) deleteClusterRole(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		ClusterRoleName string `uri:"clusterRole" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.ClusterRoleSvc.Delete(uri.ClusterRoleName))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) clusterRoleBindings(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	clusterRoleBindingList := this.ClusterRoleBindingSvc.List()

	return this.ClusterRoleBindingSvc.Paging(page, clusterRoleBindingList)
}

func (this *RBACCtl) addUserToClusterRoleBinding(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		ClusterRoleBindingName string `uri:"clusterRoleBinding" binding:"required"`
	}{}
	subject := &rbacV1.Subject{}

	athena.Error(ctx.BindUri(uri))
	athena.Error(ctx.BindJSON(subject))
	athena.Error(this.ClusterRoleBindingSvc.AddUser(uri.ClusterRoleBindingName, subject))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) removeUserFromClusterRoleBinding(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		ClusterRoleBindingName string `uri:"clusterRoleBinding" binding:"required"`
	}{}
	subject := &rbacV1.Subject{}

	athena.Error(ctx.BindUri(uri))
	athena.Error(ctx.BindJSON(subject))
	athena.Error(this.ClusterRoleBindingSvc.RemoveUser(uri.ClusterRoleBindingName, subject))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) createClusterRoleBinding(ctx *gin.Context) any {
	clusterRoleBinding := &rbacV1.ClusterRoleBinding{}
	athena.Error(ctx.BindJSON(clusterRoleBinding))
	athena.Error(this.ClusterRoleBindingSvc.Create(clusterRoleBinding))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return clusterRoleBinding
}

func (this *RBACCtl) deleteClusterRoleBinding(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		ClusterRoleBindingName string `uri:"clusterRoleBinding" binding:"required"`
	}{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.ClusterRoleBindingSvc.Delete(uri.ClusterRoleBindingName))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) userAccounts(ctx *gin.Context) any {
	page := athena.NewPageWithCtx(ctx)
	userAccountList, err := this.UserAccountService.List()
	athena.Error(err)

	return this.UserAccountService.Paging(page, userAccountList)
}

func (this *RBACCtl) createUserAccount(ctx *gin.Context) any {
	reqData := &struct {
		CN string `json:"cn" binding:"required,min=2"`
		O  string `json:"o"`
	}{}
	athena.Error(ctx.BindJSON(reqData))
	athena.Error(this.UserAccountService.Create(reqData.CN, reqData.O))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return reqData
}

func (this *RBACCtl) deleteUserAccount(ctx *gin.Context) (v athena.Void) {
	uri := &struct {
		CN string `uri:"cn" binding:"required,min=2"`
	}{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.UserAccountService.Delete(uri.CN))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *RBACCtl) userAccountKubeconfig(ctx *gin.Context) any {
	uri := &struct {
		CN string `uri:"cn" binding:"required,min=2"`
	}{}
	athena.Error(ctx.BindUri(uri))
	content := athena.Unwrap(this.UserAccountService.Kubeconfig(uri.CN)).([]byte)

	return string(content)
}

func (this *RBACCtl) Build(athena *athena.Athena) {
	// role??????
	athena.Handle("GET", "/roles", this.roles)
	// ????????????role
	athena.Handle("GET", "/roles/all", this.rolesAll)
	// ??????role
	athena.Handle("POST", "/role", this.createRole)
	// ??????role
	athena.Handle("GET", "/role/:ns/:role", this.showRole)
	// ??????role
	athena.Handle("PUT", "/role/:ns/:role", this.updateRole)
	// ??????role
	athena.Handle("DELETE", "/role/:ns/:role", this.deleteRole)
	// roleBinding??????
	athena.Handle("GET", "/roleBindings", this.roleBindings)
	// ??????roleBinding
	athena.Handle("POST", "/roleBinding", this.createRoleBinding)
	// ??????roleBinding
	athena.Handle("DELETE", "/roleBinding/:ns/:roleBinding", this.deleteRoleBinding)
	// roleBinding??????user
	athena.Handle("PATCH", "/roleBinding/:ns/:roleBinding/user", this.addUserToRoleBinding)
	// roleBinding??????user
	athena.Handle("PATCH", "/roleBinding/:ns/:roleBinding/user/remove", this.removeUserFromRoleBinding)
	// serviceAccount??????
	athena.Handle("GET", "/serviceAccounts", this.serviceAccounts)
	// ??????serviceAccount
	athena.Handle("DELETE", "/serviceAccount/:ns/:serviceAccount", this.deleteServiceAccount)
	// clusterRole??????
	athena.Handle("GET", "/clusterRoles", this.clusterRoles)
	// ????????????clusterRole
	athena.Handle("GET", "/clusterRoles/all", this.clusterRolesAll)
	// ??????clusterRole
	athena.Handle("GET", "/clusterRole/:clusterRole", this.showClusterRole)
	// ??????clusterRole
	athena.Handle("POST", "/clusterRole", this.createClusterRole)
	// ??????clusterRole
	athena.Handle("PUT", "/clusterRole/:clusterRole", this.updateClusterRole)
	// ??????clusterRole
	athena.Handle("DELETE", "/clusterRole/:clusterRole", this.deleteClusterRole)
	// clusterRoleBinding??????
	athena.Handle("GET", "/clusterRoleBindings", this.clusterRoleBindings)
	// clusterRoleBinding??????user
	athena.Handle("PATCH", "/clusterRoleBinding/:clusterRoleBinding/user", this.addUserToClusterRoleBinding)
	// clusterRoleBinding??????user
	athena.Handle("PATCH", "/clusterRoleBinding/:clusterRoleBinding/user/remove", this.removeUserFromClusterRoleBinding)
	// ??????clusterRoleBinding
	athena.Handle("POST", "/clusterRoleBinding", this.createClusterRoleBinding)
	// ??????clusterRoleBinding
	athena.Handle("DELETE", "/clusterRoleBinding/:clusterRoleBinding", this.deleteClusterRoleBinding)
	// userAccount??????
	athena.Handle("GET", "/userAccounts", this.userAccounts)
	// ??????userAccount
	athena.Handle("POST", "/userAccount", this.createUserAccount)
	// ??????userAccount
	athena.Handle("DELETE", "/userAccount/:cn", this.deleteUserAccount)
	// ??????userAccount???kubeconfig
	athena.Handle("GET", "/userAccount/:cn/kubeconfig", this.userAccountKubeconfig)
}
