package rbac

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RoleService @Service
type RoleService struct {
	RoleMap   *RoleMap              `inject:"-"`
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// ListByNs 根据ns获取role
func (this *RoleService) ListByNs(ns string) (ret []RoleListModel) {
	roleList := this.RoleMap.ListByNs(ns)
	ret = make([]RoleListModel, len(roleList))

	for i, role := range roleList {
		ret[i] = RoleListModel{
			Name:      role.Name,
			Namespace: role.Namespace,
			CreatedAt: role.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页role切片
func (this *RoleService) Paging(page *athena.Page, roleList []RoleListModel) athena.Collection {
	count := len(roleList)
	iRoleList := make([]any, count)
	for i, role := range iRoleList {
		iRoleList[i] = role
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iRoleList)
	collection := athena.NewCollection(roleList[start:end], page)
	return *collection
}

// Create 创建role
func (this *RoleService) Create(role *rbacV1.Role) error {
	role.Kind = "Role"
	role.APIVersion = "rbac.authorization.k8s.io/v1"

	_, err := this.K8sClient.RbacV1().Roles(role.Namespace).
		Create(context.Background(), role, metaV1.CreateOptions{})

	return err
}

// Delete 删除role
func (this *RoleService) Delete(ns string, roleName string) error {
	return this.K8sClient.RbacV1().Roles(ns).
		Delete(context.Background(), roleName, metaV1.DeleteOptions{})
}

// Show 查看role
func (this *RoleService) Show(ns, roleName string) RoleDetailModel {
	role := this.RoleMap.Find(ns, roleName)

	return RoleDetailModel{
		Name:      role.Name,
		Namespace: role.Namespace,
		Rules:     role.Rules,
		CreatedAt: role.CreationTimestamp.Format(athena.DateTimeFormat),
	}
}

// Update 编辑role
func (this *RoleService) Update(ns, roleName string, rules []rbacV1.PolicyRule) error {
	role := this.RoleMap.Find(ns, roleName)
	role.Rules = rules
	_, err := this.K8sClient.RbacV1().Roles(ns).
		Update(context.Background(), role, metaV1.UpdateOptions{})

	return err
}
