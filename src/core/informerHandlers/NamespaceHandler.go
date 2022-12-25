package informerHandlers

import (
	coreV1 "k8s.io/api/core/v1"
	"pixelk8/src/core/maps"
)

type NamespaceHandler struct {
	NsMap *maps.NamespaceMap `inject:"-"`
}

func (this *NamespaceHandler) OnAdd(obj interface{}) {
	this.NsMap.Add(obj.(*coreV1.Namespace))
}

func (this *NamespaceHandler) OnUpdate(oldObj, newObj interface{}) {
	this.NsMap.Update(newObj.(*coreV1.Namespace))
}

func (this *NamespaceHandler) OnDelete(obj interface{}) {
	this.NsMap.Delete(obj.(*coreV1.Namespace))
}
