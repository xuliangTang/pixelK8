package maps

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type ServiceMap struct {
	data sync.Map // key:namespace(string), value:[]*coreV1.Service
}

// Add 添加
func (this *ServiceMap) Add(service *coreV1.Service) {
	if svcList, ok := this.data.Load(service.Namespace); ok {
		svcList = append(svcList.([]*coreV1.Service), service)
		this.data.Store(service.Namespace, svcList)
	} else {
		this.data.Store(service.Namespace, []*coreV1.Service{service})
	}
}

// Update 更新
func (this *ServiceMap) Update(service *coreV1.Service) error {
	if svcList, ok := this.data.Load(service.Namespace); ok {
		svcList := svcList.([]*coreV1.Service)
		for i, svc := range svcList {
			if svc.Name == service.Name {
				svcList[i] = service
				break
			}
		}
		return nil
	}

	return fmt.Errorf("service [%s] not found", service.Name)
}

// Delete 删除
func (this *ServiceMap) Delete(service *coreV1.Service) {
	if svcList, ok := this.data.Load(service.Namespace); ok {
		svcList := svcList.([]*coreV1.Service)
		for i, svc := range svcList {
			if svc.Name == service.Name {
				newSvcList := append(svcList[:i], svcList[i+1:]...)
				this.data.Store(service.Namespace, newSvcList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的service列表
func (this *ServiceMap) ListByNs(ns string) []*coreV1.Service {
	if svcList, ok := this.data.Load(ns); ok {
		ret := svcList.([]*coreV1.Service)
		sort.Sort(coreV1Service(ret))
		return ret
	}
	return []*coreV1.Service{}
}
