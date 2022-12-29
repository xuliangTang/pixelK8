package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/properties"
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

func (this *IngressCtl) Build(athena *athena.Athena) {
	// ingress列表
	athena.Handle("GET", "/ingress", this.ingress)
}
