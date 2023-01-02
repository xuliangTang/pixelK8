package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/properties"
	"pixelk8/src/services"
)

// ConfigmapCtl @Controller
type ConfigmapCtl struct {
	CmSvc *services.ConfigmapService `inject:"-"`
}

func NewConfigmapCtl() *ConfigmapCtl {
	return &ConfigmapCtl{}
}

func (this *ConfigmapCtl) configmaps(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	cmList := this.CmSvc.ListByNs(ns)

	return this.CmSvc.Paging(page, cmList)
}

func (this *ConfigmapCtl) Build(athena *athena.Athena) {
	// configmap列表
	athena.Handle("GET", "/configmaps", this.configmaps)
}
