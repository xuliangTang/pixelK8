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

type NodeHandler struct {
	NodeMap *maps.NodeMap         `inject:"-"`
	NodeSvc *services.NodeService `inject:"-"`
}

func (this *NodeHandler) OnAdd(obj interface{}) {
	node := obj.(*coreV1.Node)
	this.NodeMap.Add(node)

	// 通知ws客户端
	nodeList := this.NodeSvc.List()
	page := athena.NewPage(1, 10)
	ws.ClientMap.SendAll(dto.WS{
		Type: "nodes",
		Result: gin.H{
			"ns":   "",
			"data": this.NodeSvc.Paging(page, nodeList),
		},
	})
}

func (this *NodeHandler) OnUpdate(oldObj, newObj interface{}) {
	node := newObj.(*coreV1.Node)
	this.NodeMap.Update(node)

	nodeList := this.NodeSvc.List()
	page := athena.NewPage(1, 10)
	ws.ClientMap.SendAll(dto.WS{
		Type: "nodes",
		Result: gin.H{
			"ns":   "",
			"data": this.NodeSvc.Paging(page, nodeList),
		},
	})
}

func (this *NodeHandler) OnDelete(obj interface{}) {
	node := obj.(*coreV1.Node)
	this.NodeMap.Delete(node)

	nodeList := this.NodeSvc.List()
	page := athena.NewPage(1, 10)
	ws.ClientMap.SendAll(dto.WS{
		Type: "nodes",
		Result: gin.H{
			"ns":   "",
			"data": this.NodeSvc.Paging(page, nodeList),
		},
	})
}
