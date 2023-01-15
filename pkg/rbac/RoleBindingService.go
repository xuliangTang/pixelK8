package rbac

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

// RoleBindingService @Service
type RoleBindingService struct {
	RoleBindingMap *RoleBindingMap `inject:"-"`
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
