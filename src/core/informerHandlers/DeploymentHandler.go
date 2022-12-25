package informerHandlers

import (
	appsV1 "k8s.io/api/apps/v1"
	"log"
	"pixelk8/src/core/maps"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

type DeploymentHandler struct {
	DeploymentMap *maps.DeploymentMap         `inject:"-"`
	DeploymentSvc *services.DeploymentService `inject:"-"`
}

func (this *DeploymentHandler) OnAdd(obj interface{}) {
	this.DeploymentMap.Add(obj.(*appsV1.Deployment))
	ws.ClientMap.SendDeploymentList(this.DeploymentSvc.ListByNs(obj.(*appsV1.Deployment).Namespace))
}

func (this *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.DeploymentMap.Update(newObj.(*appsV1.Deployment))
	if err != nil {
		log.Println(err)
	} else {
		ws.ClientMap.SendDeploymentList(this.DeploymentSvc.ListByNs(newObj.(*appsV1.Deployment).Namespace))
	}
}

func (this *DeploymentHandler) OnDelete(obj interface{}) {
	this.DeploymentMap.Delete(obj.(*appsV1.Deployment))
	ws.ClientMap.SendDeploymentList(this.DeploymentSvc.ListByNs(obj.(*appsV1.Deployment).Namespace))
}
