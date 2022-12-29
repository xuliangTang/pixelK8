package services

import (
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
)

// ServiceService @Service
type ServiceService struct {
	SvcMap *maps.ServiceMap `inject:"-"`
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
			CreatedAt: svc.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}
