package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

// ClusterRoleService @Service
type ClusterRoleService struct {
	ClusterRoleMap *ClusterRoleMap `inject:"-"`
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
