package informerHandlers

import (
	appsV1 "k8s.io/api/apps/v1"
	"log"
	"pixelk8/src/core/maps"
)

type DeploymentHandler struct {
	DeploymentMap *maps.DeploymentMap `inject:"-"`
}

func (this *DeploymentHandler) OnAdd(obj interface{}) {
	this.DeploymentMap.Add(obj.(*appsV1.Deployment))
}

func (this *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.DeploymentMap.Update(newObj.(*appsV1.Deployment))
	if err != nil {
		log.Println(err)
	}
}

func (this *DeploymentHandler) OnDelete(obj interface{}) {
	this.DeploymentMap.Delete(obj.(*appsV1.Deployment))
}
