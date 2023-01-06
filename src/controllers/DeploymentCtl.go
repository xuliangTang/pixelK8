package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"net/http"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
)

// DeploymentCtl @controller
type DeploymentCtl struct {
	DeploymentService *services.DeploymentService `inject:"-"`
	PodService        *services.PodService        `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl) deployments(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	depList := this.DeploymentService.ListByNs(ns)

	return this.DeploymentService.Paging(page, depList)
}

func (this *DeploymentCtl) deploymentPods(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	var uri requests.DeploymentUri
	athena.Error(ctx.BindUri(&uri))
	podList := this.PodService.ListByDeployment(ns, uri.Deployment)

	return podList
}

func (this *DeploymentCtl) deleteDeployment(ctx *gin.Context) (v athena.Void) {
	uri := &requests.DeleteDeploymentUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.DeploymentService.Delete(uri))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *DeploymentCtl) Build(athena *athena.Athena) {
	// Deployment 列表
	athena.Handle("GET", "/deployments", this.deployments)
	// Deployment Pod 列表
	athena.Handle("GET", "/deployment/:deployment/pods", this.deploymentPods)
	// 删除 Deployment
	athena.Handle("DELETE", "/deployment/:ns/:deployment", this.deleteDeployment)
}
