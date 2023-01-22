package rbac

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceAccountService @Service
type ServiceAccountService struct {
	SaMap     *ServiceAccountMap    `inject:"-"`
	K8sClient *kubernetes.Clientset `inject:"-"`
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

func (this *ServiceAccountService) Delete(ns, name string) error {
	return this.K8sClient.CoreV1().ServiceAccounts(ns).
		Delete(context.Background(), name, metaV1.DeleteOptions{})
}
