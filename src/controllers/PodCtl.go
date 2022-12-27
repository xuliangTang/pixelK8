package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/properties"
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

	// pod总数和就绪数
	var count, countReady int
	iPodList := make([]any, len(podList))
	for i, pod := range podList {
		iPodList[i] = pod
		count++
		if pod.IsReady {
			countReady++
		}
	}

	page.Extend = gin.H{"count": count, "count_ready": countReady}
	// 分页
	start, end := page.SlicePage(iPodList)
	collection := athena.NewCollection(podList[start:end], page)

	return *collection
}

func (this *PodCtl) Build(athena *athena.Athena) {
	// 获取pod列表
	athena.Handle("GET", "/pods", this.pods)
}
