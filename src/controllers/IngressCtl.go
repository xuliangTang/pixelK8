package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"net/http"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
)

// IngressCtl @Controller
type IngressCtl struct {
	IngSvc *services.IngressService `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (this *IngressCtl) ingress(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	ingList := this.IngSvc.ListByNs(ns)

	return this.IngSvc.Paging(page, ingList)
}

func (this *IngressCtl) createIngress(ctx *gin.Context) (athena.HttpCode, any) {
	req := &requests.CreateIngress{}
	athena.Error(ctx.BindJSON(req))
	athena.Error(this.IngSvc.Create(req))

	return http.StatusCreated, req
}

func (this *IngressCtl) deleteIngress(ctx *gin.Context) athena.HttpCode {
	uri := &requests.DeleteIngressUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.IngSvc.Delete(uri))

	return http.StatusNoContent
}

func (this *IngressCtl) Build(athena *athena.Athena) {
	// ingress列表
	athena.Handle("GET", "/ingress", this.ingress)
	// 创建ingress
	athena.Handle("POST", "/ingress", this.createIngress)
	// 删除ingress
	athena.Handle("DELETE", "/ingress/:ns/:ingress", this.deleteIngress)
}
