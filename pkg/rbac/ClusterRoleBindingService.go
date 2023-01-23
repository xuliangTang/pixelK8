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

// ClusterRoleBindingService @Service
type ClusterRoleBindingService struct {
	ClusterRoleBindingMap *ClusterRoleBindingMap `inject:"-"`
	K8sClient             *kubernetes.Clientset  `inject:"-"`
}

func NewClusterRoleBindingService() *ClusterRoleBindingService {
	return &ClusterRoleBindingService{}
}

// List 获取所有clusterRoleBinding
func (this *ClusterRoleBindingService) List() (ret []ClusterRoleBindingListModel) {
	clusterRoleBindingList := this.ClusterRoleBindingMap.List()

	ret = make([]ClusterRoleBindingListModel, len(clusterRoleBindingList))
	for i, clusterRoleBinding := range clusterRoleBindingList {
		ret[i] = ClusterRoleBindingListModel{
			Name:      clusterRoleBinding.Name,
			RoleRef:   fmt.Sprintf("%s / %s", clusterRoleBinding.RoleRef.Kind, clusterRoleBinding.RoleRef.Name),
			Subjects:  clusterRoleBinding.Subjects,
			CreatedAt: clusterRoleBinding.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页clusterRoleBinding切片
func (this *ClusterRoleBindingService) Paging(page *athena.Page, clusterRoleBindingList []ClusterRoleBindingListModel) athena.Collection {
	count := len(clusterRoleBindingList)
	iClusterRoleBindingList := make([]any, count)

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iClusterRoleBindingList)
	collection := athena.NewCollection(clusterRoleBindingList[start:end], page)
	return *collection
}

// AddUser 向clusterRoleBinding添加subject
func (this *ClusterRoleBindingService) AddUser(clusterRoleBindingName string, subject *rbacV1.Subject) error {
	clusterRoleBinding := *this.ClusterRoleBindingMap.Find(clusterRoleBindingName)
	clusterRoleBinding.Subjects = append(clusterRoleBinding.Subjects, *subject)

	_, err := this.K8sClient.RbacV1().ClusterRoleBindings().
		Update(context.Background(), &clusterRoleBinding, metaV1.UpdateOptions{})

	return err
}

// RemoveUser 从clusterRoleBinding中移除subject
func (this *ClusterRoleBindingService) RemoveUser(clusterRoleBindingName string, subject *rbacV1.Subject) error {
	clusterRoleBinding := *this.ClusterRoleBindingMap.Find(clusterRoleBindingName)

	for i, sub := range clusterRoleBinding.Subjects {
		if sub.Kind == subject.Kind && sub.Name == subject.Name {
			clusterRoleBinding.Subjects = append(clusterRoleBinding.Subjects[:i], clusterRoleBinding.Subjects[i+1:]...)
			break
		}
	}

	_, err := this.K8sClient.RbacV1().ClusterRoleBindings().
		Update(context.Background(), &clusterRoleBinding, metaV1.UpdateOptions{})

	return err
}

// Create 创建clusterRoleBinding
func (this *ClusterRoleBindingService) Create(clusterRoleBinding *rbacV1.ClusterRoleBinding) error {
	clusterRoleBinding.Kind = "ClusterRoleBinding"
	clusterRoleBinding.APIVersion = "rbac.authorization.k8s.io/v1"

	_, err := this.K8sClient.RbacV1().ClusterRoleBindings().
		Create(context.Background(), clusterRoleBinding, metaV1.CreateOptions{})

	return err
}

// Delete 删除clusterRoleBinding
func (this *ClusterRoleBindingService) Delete(clusterRoleBindingName string) error {
	return this.K8sClient.RbacV1().ClusterRoleBindings().
		Delete(context.Background(), clusterRoleBindingName, metaV1.DeleteOptions{})
}
