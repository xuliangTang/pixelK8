package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

// RoleService @Service
type RoleService struct {
	RoleMap *RoleMap `inject:"-"`
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
