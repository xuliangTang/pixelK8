package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
)

// PodCtl @Controller
type PodCtl struct {
	PodService *services.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (this *PodCtl) pods(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	podList := this.PodService.ListByNs(ns)

	return this.PodService.Paging(page, podList)
}

func (this *PodCtl) podAllContainers(ctx *gin.Context) any {
	uri := &requests.PodAllContainersUri{}
	athena.Error(ctx.BindUri(uri))

	return this.PodService.GetPodContainers(uri)
}

func (this *PodCtl) Build(athena *athena.Athena) {
	// 获取pod列表
	athena.Handle("GET", "/pods", this.pods)
	// 获取pod所有容器
	athena.Handle("GET", "/pod/:ns/:pod/containers", this.podAllContainers)
}
