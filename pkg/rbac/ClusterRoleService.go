package rbac

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClusterRoleService @Service
type ClusterRoleService struct {
	ClusterRoleMap *ClusterRoleMap       `inject:"-"`
	K8sClient      *kubernetes.Clientset `inject:"-"`
}

func NewClusterRoleService() *ClusterRoleService {
	return &ClusterRoleService{}
}

// List 获取clusterRole列表
func (this *ClusterRoleService) List() (ret []ClusterRoleListModel) {
	clusterRoleList := this.ClusterRoleMap.List()

	ret = make([]ClusterRoleListModel, len(clusterRoleList))
	for i, clusterRole := range clusterRoleList {
		ret[i] = ClusterRoleListModel{
			Name:      clusterRole.Name,
			CreatedAt: clusterRole.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页clusterRole切片
func (this *ClusterRoleService) Paging(page *athena.Page, clusterRoleList []ClusterRoleListModel) athena.Collection {
	count := len(clusterRoleList)
	iClusterRoleList := make([]any, count)
	for i, clusterRole := range clusterRoleList {
		iClusterRoleList[i] = clusterRole
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iClusterRoleList)
	collection := athena.NewCollection(clusterRoleList[start:end], page)
	return *collection
}

// Show 查看clusterRole
func (this *ClusterRoleService) Show(name string) ClusterRoleDetailModel {
	clusterRole := this.ClusterRoleMap.Find(name)

	return ClusterRoleDetailModel{
		Name:        clusterRole.Name,
		Rules:       clusterRole.Rules,
		Labels:      clusterRole.Labels,
		Annotations: clusterRole.Annotations,
		CreatedAt:   clusterRole.CreationTimestamp.Format(athena.DateTimeFormat),
	}
}

func (this *ClusterRoleService) Create(clusterRole *rbacV1.ClusterRole) error {
	clusterRole.Kind = "ClusterRole"
	clusterRole.APIVersion = "ClusterRole"

	_, err := this.K8sClient.RbacV1().ClusterRoles().
		Create(context.Background(), clusterRole, metaV1.CreateOptions{})

	return err
}

func (this *ClusterRoleService) Update(name string, rules []rbacV1.PolicyRule) error {
	clusterRole := *this.ClusterRoleMap.Find(name)
	clusterRole.Rules = rules

	_, err := this.K8sClient.RbacV1().ClusterRoles().
		Update(context.Background(), &clusterRole, metaV1.UpdateOptions{})

	return err
}
