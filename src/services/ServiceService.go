package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	coreV1 "k8s.io/api/core/v1"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
)

// ServiceService @Service
type ServiceService struct {
	SvcMap    *maps.ServiceMap `inject:"-"`
	CommonSvc *CommonService   `inject:"-"`
}

func NewServiceService() *ServiceService {
	return &ServiceService{}
}

// ListByNs 获取ns下的services
func (this *ServiceService) ListByNs(ns string) (ret []*dto.ServiceList) {
	svcList := this.SvcMap.ListByNs(ns)

	ret = make([]*dto.ServiceList, len(svcList))
	for i, svc := range svcList {
		ret[i] = &dto.ServiceList{
			Name:      svc.Name,
			Namespace: svc.Namespace,
			Selector:  this.CommonSvc.SelectorToStrings(svc.Spec.Selector),
			Type:      string(svc.Spec.Type),
			Target:    this.getPortsTarget(svc),
			CreatedAt: svc.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页service切片
func (this *ServiceService) Paging(page *athena.Page, svcList []*dto.ServiceList) athena.Collection {
	var count int
	iSvcList := make([]any, len(svcList))
	for i, svc := range iSvcList {
		count++
		iSvcList[i] = svc
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iSvcList)
	collection := athena.NewCollection(svcList[start:end], page)
	return *collection
}

// 拼接service target
func (this *ServiceService) getPortsTarget(service *coreV1.Service) (ret []string) {
	ret = make([]string, 0)
	for _, p := range service.Spec.Ports {
		ret = append(ret, fmt.Sprintf("%s %d/%s", p.Name, p.Port, p.Protocol))
	}
	return
}
