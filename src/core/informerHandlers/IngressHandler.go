package informerHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	networkingV1 "k8s.io/api/networking/v1"
	"log"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

type IngressHandler struct {
	IngMap *maps.IngressMap         `inject:"-"`
	IngSvc *services.IngressService `inject:"-"`
}

func (this *IngressHandler) OnAdd(obj interface{}) {
	ing := obj.(*networkingV1.Ingress)
	this.IngMap.Add(ing)

	// 通知ws客户端
	ingList := this.IngSvc.ListByNs(ing.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "ingress",
		Result: gin.H{
			"ns":   ing.Namespace,
			"data": this.IngSvc.Paging(page, ingList),
		},
	})
}
func (this *IngressHandler) OnUpdate(oldObj, newObj interface{}) {
	ing := newObj.(*networkingV1.Ingress)
	err := this.IngMap.Update(ing)
	if err != nil {
		log.Println(err)
	} else {
		ingList := this.IngSvc.ListByNs(ing.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "ingress",
			Result: gin.H{
				"ns":   ing.Namespace,
				"data": this.IngSvc.Paging(page, ingList),
			},
		})
	}
}
func (this *IngressHandler) OnDelete(obj interface{}) {
	ing := obj.(*networkingV1.Ingress)
	this.IngMap.Delete(ing)

	ingList := this.IngSvc.ListByNs(ing.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "ingress",
		Result: gin.H{
			"ns":   ing.Namespace,
			"data": this.IngSvc.Paging(page, ingList),
		},
	})
}
