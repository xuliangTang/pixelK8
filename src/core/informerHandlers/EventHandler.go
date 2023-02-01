package informerHandlers

import (
	coreV1 "k8s.io/api/core/v1"
	"pixelk8/src/core/maps"
)

type EventHandler struct {
	EventMap *maps.EventMap `inject:"-"`
}

func (this *EventHandler) OnAdd(obj interface{}) {
	if event, ok := obj.(*coreV1.Event); ok {
		this.EventMap.Add(event)
	}
}

func (this *EventHandler) OnUpdate(oldObj, newObj interface{}) {
	if event, ok := newObj.(*coreV1.Event); ok {
		this.EventMap.Add(event)
	}
}

func (this *EventHandler) OnDelete(obj interface{}) {
	if event, ok := obj.(*coreV1.Event); ok {
		this.EventMap.Delete(event)
	}
}
