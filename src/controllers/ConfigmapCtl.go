package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"net/http"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
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

func (this *ConfigmapCtl) createConfigmap(ctx *gin.Context) (athena.HttpCode, any) {
	req := &requests.CreateConfigmap{}
	athena.Error(ctx.BindJSON(req))
	athena.Error(this.CmSvc.Create(req))

	return http.StatusCreated, req
}

func (this *ConfigmapCtl) showConfigmap(ctx *gin.Context) any {
	uri := &requests.ShowConfigmapUri{}
	athena.Error(ctx.BindUri(uri))

	return this.CmSvc.Show(uri)
}

func (this *ConfigmapCtl) Build(athena *athena.Athena) {
	// configmap列表
	athena.Handle("GET", "/configmaps", this.configmaps)
	// 创建configmap
	athena.Handle("POST", "/configmap", this.createConfigmap)
	// 查看configmap
	athena.Handle("GET", "/configmap/:ns/:configmap", this.showConfigmap)
}
