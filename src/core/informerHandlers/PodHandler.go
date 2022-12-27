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

type PodHandler struct {
	PodMap *maps.PodMap         `inject:"-"`
	PodSvc *services.PodService `inject:"-"`
}

func (this *PodHandler) OnAdd(obj interface{}) {
	pod := obj.(*coreV1.Pod)
	this.PodMap.Add(pod)

	// 通知ws客户端
	podList := this.PodSvc.ListByNs(pod.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "pods",
		Result: gin.H{
			"ns":   pod.Namespace,
			"data": this.PodSvc.Paging(page, podList),
		},
	})
}
func (this *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	pod := newObj.(*coreV1.Pod)
	err := this.PodMap.Update(pod)
	if err != nil {
		log.Println(err)
	}

	podList := this.PodSvc.ListByNs(pod.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "pods",
		Result: gin.H{
			"ns":   pod.Namespace,
			"data": this.PodSvc.Paging(page, podList),
		},
	})
}
func (this *PodHandler) OnDelete(obj interface{}) {
	pod := obj.(*coreV1.Pod)
	this.PodMap.Delete(pod)

	podList := this.PodSvc.ListByNs(pod.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "pods",
		Result: gin.H{
			"ns":   pod.Namespace,
			"data": this.PodSvc.Paging(page, podList),
		},
	})
}
