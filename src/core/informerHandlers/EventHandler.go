package informerHandlers

import (
	coreV1 "k8s.io/api/core/v1"
	"pixelk8/src/core/maps"
)

type EventHandler struct {
	EventMap *maps.EventMap `inject:"-"`
}

func (this *EventHandler) OnAdd(obj interface{}) {
	this.EventMap.Add(obj.(*coreV1.Event))
}
func (this *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	this.EventMap.Add(newObj.(*coreV1.Event))
}
func (this *EventHandler) OnDelete(obj interface{}) {
	this.EventMap.Delete(obj.(*coreV1.Event))
}
