package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/services"
)

// NodeCtl @Controller
type NodeCtl struct {
	NodeService *services.NodeService `inject:"-"`
}

func NewNodeCtl() *NodeCtl {
	return &NodeCtl{}
}

func (this *NodeCtl) nodes(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	nodeList := this.NodeService.List()

	return this.NodeService.Paging(page, nodeList)
}

func (this *NodeCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/nodes", this.nodes)
}
