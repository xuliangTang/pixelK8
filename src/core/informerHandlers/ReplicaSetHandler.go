package informerHandlers

import (
	appsV1 "k8s.io/api/apps/v1"
	"log"
	"pixelk8/src/core/maps"
)

type ReplicaSetHandler struct {
	ReplicaSetMap *maps.ReplicaSetMap `inject:"-"`
}

func (this *ReplicaSetHandler) OnAdd(obj interface{}) {
	this.ReplicaSetMap.Add(obj.(*appsV1.ReplicaSet))
}

func (this *ReplicaSetHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.ReplicaSetMap.Update(newObj.(*appsV1.ReplicaSet))
	if err != nil {
		log.Println(err)
	}
}

func (this *ReplicaSetHandler) OnDelete(obj interface{}) {
	this.ReplicaSetMap.Delete(obj.(*appsV1.ReplicaSet))
}
