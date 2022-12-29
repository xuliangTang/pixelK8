package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
)

// IngressService @Service
type IngressService struct {
	IngressMap *maps.IngressMap `inject:"-"`
}

func NewIngressService() *IngressService {
	return &IngressService{}
}

// ListByNs 根据ns获取ingress列表
func (this *IngressService) ListByNs(ns string) (ret []*dto.IngressList) {
	ingList := this.IngressMap.ListByNs(ns)

	ret = make([]*dto.IngressList, len(ingList))
	for i, ing := range ingList {
		ret[i] = &dto.IngressList{
			Name:      ing.Name,
			CreatedAt: ing.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页ingress切片
func (this *IngressService) Paging(page *athena.Page, ingList []*dto.IngressList) athena.Collection {
	var count int
	iIngList := make([]any, len(ingList))
	for i, ing := range ingList {
		count++
		iIngList[i] = ing
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iIngList)
	collection := athena.NewCollection(ingList[start:end], page)
	return *collection
}
