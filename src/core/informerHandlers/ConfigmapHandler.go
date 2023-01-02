package informerHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	coreV1 "k8s.io/api/core/v1"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

type ConfigmapHandler struct {
	CmMap *maps.ConfigmapMap         `inject:"-"`
	CmSvc *services.ConfigmapService `inject:"-"`
}

func (this *ConfigmapHandler) OnAdd(obj interface{}) {
	cm := obj.(*coreV1.ConfigMap)
	this.CmMap.Add(cm)

	// 通知ws客户端
	cmList := this.CmSvc.ListByNs(cm.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "configmaps",
		Result: gin.H{
			"ns":   cm.Namespace,
			"data": this.CmSvc.Paging(page, cmList),
		},
	})
}

func (this *ConfigmapHandler) OnUpdate(oldObj, newObj interface{}) {
	cm := newObj.(*coreV1.ConfigMap)
	if this.CmMap.Update(cm) {
		cmList := this.CmSvc.ListByNs(cm.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "configmaps",
			Result: gin.H{
				"ns":   cm.Namespace,
				"data": this.CmSvc.Paging(page, cmList),
			},
		})
	}
}

func (this *ConfigmapHandler) OnDelete(obj interface{}) {
	cm := obj.(*coreV1.ConfigMap)
	this.CmMap.Delete(cm)

	cmList := this.CmSvc.ListByNs(cm.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "configmaps",
		Result: gin.H{
			"ns":   cm.Namespace,
			"data": this.CmSvc.Paging(page, cmList),
		},
	})
}
