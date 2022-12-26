package informerHandlers

import (
	"github.com/gin-gonic/gin"
	appsV1 "k8s.io/api/apps/v1"
	"log"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

type DeploymentHandler struct {
	DeploymentMap *maps.DeploymentMap         `inject:"-"`
	DeploymentSvc *services.DeploymentService `inject:"-"`
}

func (this *DeploymentHandler) OnAdd(obj interface{}) {
	dep := obj.(*appsV1.Deployment)
	this.DeploymentMap.Add(dep)

	// 通知ws客户端
	ws.ClientMap.SendAll(dto.WS{
		Type: "deployments",
		Result: gin.H{
			"ns":   dep.Namespace,
			"data": this.DeploymentSvc.ListByNs(dep.Namespace),
		},
	})
}

func (this *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	dep := newObj.(*appsV1.Deployment)
	err := this.DeploymentMap.Update(dep)
	if err != nil {
		log.Println(err)
	} else {
		ws.ClientMap.SendAll(dto.WS{
			Type: "deployments",
			Result: gin.H{
				"ns":   dep.Namespace,
				"data": this.DeploymentSvc.ListByNs(dep.Namespace),
			},
		})
	}
}

func (this *DeploymentHandler) OnDelete(obj interface{}) {
	dep := obj.(*appsV1.Deployment)
	this.DeploymentMap.Delete(dep)
	ws.ClientMap.SendAll(dto.WS{
		Type: "deployments",
		Result: gin.H{
			"ns":   dep.Namespace,
			"data": this.DeploymentSvc.ListByNs(dep.Namespace),
		},
	})
}
