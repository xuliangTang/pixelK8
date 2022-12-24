package maps

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"sync"
)

type EventMap struct {
	Data sync.Map // key:namespace_kind_name value: *v1.Event
}

func (this *EventMap) GetKey(event *coreV1.Event) string {
	key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
	return key
}

// Add 添加
func (this *EventMap) Add(event *coreV1.Event) {
	this.Data.Store(this.GetKey(event), event)
}

// Delete 删除
func (this *EventMap) Delete(event *coreV1.Event) {
	this.Data.Delete(this.GetKey(event))
}

func (this *EventMap) GetMessage(ns string, kind string, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	if v, ok := this.Data.Load(key); ok {
		return v.(*coreV1.Event).Message
	}

	return ""
}
