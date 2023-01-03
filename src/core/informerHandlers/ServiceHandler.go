package informerHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	coreV1 "k8s.io/api/core/v1"
	"log"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

type ServiceHandler struct {
	SvcMap     *maps.ServiceMap         `inject:"-"`
	SvcService *services.ServiceService `inject:"-"`
}

func (this *ServiceHandler) OnAdd(obj interface{}) {
	svc := obj.(*coreV1.Service)
	this.SvcMap.Add(svc)

	// 通知ws客户端
	svcList := this.SvcService.ListByNs(svc.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "services",
		Result: gin.H{
			"ns":   svc.Namespace,
			"data": this.SvcService.Paging(page, svcList),
		},
	})
}

func (this *ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	svc := newObj.(*coreV1.Service)
	err := this.SvcMap.Update(svc)
	if err != nil {
		log.Println(err)
	} else {
		svcList := this.SvcService.ListByNs(svc.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "services",
			Result: gin.H{
				"ns":   svc.Namespace,
				"data": this.SvcService.Paging(page, svcList),
			},
		})
	}
}

func (this *ServiceHandler) OnDelete(obj interface{}) {
	svc := obj.(*coreV1.Service)
	this.SvcMap.Delete(svc)

	svcList := this.SvcService.ListByNs(svc.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "services",
		Result: gin.H{
			"ns":   svc.Namespace,
			"data": this.SvcService.Paging(page, svcList),
		},
	})
}
