package rbac

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RoleBindingService @Service
type RoleBindingService struct {
	RoleBindingMap *RoleBindingMap       `inject:"-"`
	K8sClient      *kubernetes.Clientset `inject:"-"`
}

func NewRoleBindingService() *RoleBindingService {
	return &RoleBindingService{}
}

// ListByNs 获取ns下的roleBindings
func (this *RoleBindingService) ListByNs(ns string) (ret []RoleBindingListModel) {
	roleBindingList := this.RoleBindingMap.ListByNs(ns)

	ret = make([]RoleBindingListModel, len(roleBindingList))
	for i, roleBinding := range roleBindingList {
		ret[i] = RoleBindingListModel{
			Name:      roleBinding.Name,
			Namespace: roleBinding.Namespace,
			RoleRef:   fmt.Sprintf("%s / %s", roleBinding.RoleRef.Kind, roleBinding.RoleRef.Name),
			Subjects:  roleBinding.Subjects,
			CreatedAt: roleBinding.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页role切片
func (this *RoleBindingService) Paging(page *athena.Page, roleBindingList []RoleBindingListModel) athena.Collection {
	count := len(roleBindingList)
	iRoleBindingList := make([]any, count)
	for i, roleBinding := range roleBindingList {
		iRoleBindingList[i] = roleBinding
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iRoleBindingList)
	collection := athena.NewCollection(roleBindingList[start:end], page)
	return *collection
}

// Create 创建roleBinding
func (this *RoleBindingService) Create(roleBinding *rbacV1.RoleBinding) error {
	roleBinding.Kind = "RoleBinding"
	roleBinding.APIVersion = "rbac.authorization.k8s.io/v1"

	_, err := this.K8sClient.RbacV1().RoleBindings(roleBinding.Namespace).
		Create(context.Background(), roleBinding, metaV1.CreateOptions{})

	return err
}

// AddUser 向roleBinding添加user
func (this *RoleBindingService) AddUser(ns, roleBindingName string, subject *rbacV1.Subject) error {
	roleBinding := *this.RoleBindingMap.Find(ns, roleBindingName) // 取指针的值拷贝，否则会修改原对象
	roleBinding.Subjects = append(roleBinding.Subjects, *subject)

	_, err := this.K8sClient.RbacV1().RoleBindings(ns).
		Update(context.Background(), &roleBinding, metaV1.UpdateOptions{})

	return err
}

// RemoveUser 从roleBinding移除user
func (this *RoleBindingService) RemoveUser(ns, roleBindingName string, subject *rbacV1.Subject) error {
	roleBinding := *this.RoleBindingMap.Find(ns, roleBindingName)

	for i, sub := range roleBinding.Subjects {
		if sub.Kind == subject.Kind && sub.Name == subject.Name {
			roleBinding.Subjects = append(roleBinding.Subjects[:i], roleBinding.Subjects[i+1:]...)
			break
		}
	}

	_, err := this.K8sClient.RbacV1().RoleBindings(ns).
		Update(context.Background(), &roleBinding, metaV1.UpdateOptions{})

	return err
}

// Delete 删除roleBinding
func (this *RoleBindingService) Delete(ns, roleBindingName string) error {
	return this.K8sClient.RbacV1().RoleBindings(ns).
		Delete(context.Background(), roleBindingName, metaV1.DeleteOptions{})
}
