package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

// ServiceAccountService @Service
type ServiceAccountService struct {
	SaMap *ServiceAccountMap `inject:"-"`
}

func NewServiceAccountService() *ServiceAccountService {
	return &ServiceAccountService{}
}

// ListByNs 获取ns下的sa
func (this *ServiceAccountService) ListByNs(ns string) (ret []ServiceAccountListModel) {
	saList := this.SaMap.ListByNs(ns)

	ret = make([]ServiceAccountListModel, len(saList))
	for i, sa := range saList {
		ret[i] = ServiceAccountListModel{
			Name:      sa.Name,
			Namespace: sa.Namespace,
			Secrets:   sa.Secrets,
			CreatedAt: sa.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页role切片
func (this *ServiceAccountService) Paging(page *athena.Page, saList []ServiceAccountListModel) athena.Collection {
	count := len(saList)
	iSaList := make([]any, count)
	for i, sa := range saList {
		iSaList[i] = sa
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iSaList)
	collection := athena.NewCollection(saList[start:end], page)
	return *collection
}
