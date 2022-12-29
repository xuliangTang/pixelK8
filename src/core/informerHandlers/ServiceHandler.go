package informerHandlers

import (
	coreV1 "k8s.io/api/core/v1"
	"log"
	"pixelk8/src/core/maps"
)

type ServiceHandler struct {
	SvcMap *maps.ServiceMap `inject:"-"`
}

func (this *ServiceHandler) OnAdd(obj interface{}) {
	svc := obj.(*coreV1.Service)
	this.SvcMap.Add(svc)
}

func (this *ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	svc := newObj.(*coreV1.Service)
	err := this.SvcMap.Update(svc)
	if err != nil {
		log.Println(err)
	}
}

func (this *ServiceHandler) OnDelete(obj interface{}) {
	svc := obj.(*coreV1.Service)
	this.SvcMap.Delete(svc)
}
