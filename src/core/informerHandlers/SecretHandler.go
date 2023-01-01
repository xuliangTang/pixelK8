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

type SecretHandler struct {
	SecretMap *maps.SecretMap         `inject:"-"`
	SecretSvc *services.SecretService `inject:"-"`
}

func (this *SecretHandler) OnAdd(obj interface{}) {
	ing := obj.(*coreV1.Secret)
	this.SecretMap.Add(ing)

	// 通知ws客户端
	secList := this.SecretSvc.ListByNs(ing.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "secrets",
		Result: gin.H{
			"ns":   ing.Namespace,
			"data": this.SecretSvc.Paging(page, secList),
		},
	})
}

func (this *SecretHandler) OnUpdate(oldObj, newObj interface{}) {
	ing := newObj.(*coreV1.Secret)
	err := this.SecretMap.Update(ing)
	if err != nil {
		log.Println(err)
	} else {
		secList := this.SecretSvc.ListByNs(ing.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "secrets",
			Result: gin.H{
				"ns":   ing.Namespace,
				"data": this.SecretSvc.Paging(page, secList),
			},
		})
	}
}

func (this *SecretHandler) OnDelete(obj interface{}) {
	ing := obj.(*coreV1.Secret)
	this.SecretMap.Delete(ing)

	secList := this.SecretSvc.ListByNs(ing.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "secrets",
		Result: gin.H{
			"ns":   ing.Namespace,
			"data": this.SecretSvc.Paging(page, secList),
		},
	})
}
