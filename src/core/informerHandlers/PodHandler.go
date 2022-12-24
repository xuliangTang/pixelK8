package informerHandlers

import (
	coreV1 "k8s.io/api/core/v1"
	"log"
	"pixelk8/src/core/maps"
)

type PodHandler struct {
	PodMap *maps.PodMap `inject:"-"`
}

func (this *PodHandler) OnAdd(obj interface{}) {
	this.PodMap.Add(obj.(*coreV1.Pod))
}
func (this *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.PodMap.Update(newObj.(*coreV1.Pod))
	if err != nil {
		log.Println(err)
	}
}
func (this *PodHandler) OnDelete(obj interface{}) {
	this.PodMap.Delete(obj.(*coreV1.Pod))
}
