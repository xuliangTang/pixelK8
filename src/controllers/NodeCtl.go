package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/requests"
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

func (this *NodeCtl) showNode(ctx *gin.Context) any {
	uri := &requests.ShowNodeUri{}
	athena.Error(ctx.BindUri(uri))

	return this.NodeService.Show(uri)
}

func (this *NodeCtl) update(ctx *gin.Context) (v athena.Void) {
	uri := &requests.ShowNodeUri{}
	athena.Error(ctx.BindUri(uri))

	req := &requests.UpdateNode{}
	athena.Error(ctx.BindJSON(req))

	athena.Error(this.NodeService.Update(uri, req))
	return
}

func (this *NodeCtl) Build(athena *athena.Athena) {
	// 获取node列表
	athena.Handle("GET", "/nodes", this.nodes)
	// 查看node详情
	athena.Handle("GET", "/node/:node", this.showNode)
	// 更新node标签和污点
	athena.Handle("PATCH", "/node/:node", this.update)
}
