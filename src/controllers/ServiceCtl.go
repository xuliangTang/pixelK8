package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/properties"
	"pixelk8/src/services"
)

// ServiceCtl @Controller
type ServiceCtl struct {
	SvcService *services.ServiceService `inject:"-"`
}

func NewServiceCtl() *ServiceCtl {
	return &ServiceCtl{}
}

func (this *ServiceCtl) serviceAll(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	return this.SvcService.ListByNs(ns)
}

func (this *ServiceCtl) services(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	svcList := this.SvcService.ListByNs(ns)

	return this.SvcService.Paging(page, svcList)
}

func (this *ServiceCtl) Build(athena *athena.Athena) {
	// 获取所有services
	athena.Handle("GET", "/service/all", this.serviceAll)
	// service列表
	athena.Handle("GET", "/services", this.services)
}
