package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	appsV1 "k8s.io/api/apps/v1"
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

func (this *DeploymentCtl) showDeployment(ctx *gin.Context) any {
	uri := &requests.ShowDeploymentUri{}
	athena.Error(ctx.BindUri(uri))

	return this.DeploymentService.Show(uri)
}

func (this *DeploymentCtl) deploymentInfo(ctx *gin.Context) any {
	uri := &requests.ShowDeploymentUri{}
	athena.Error(ctx.BindUri(uri))

	return this.DeploymentService.Info(uri)
}

func (this *DeploymentCtl) deploymentPods(ctx *gin.Context) any {
	var uri requests.DeploymentUri
	athena.Error(ctx.BindUri(&uri))
	podList := this.PodService.ListByDeployment(uri.Namespace, uri.Deployment)

	return podList
}

func (this *DeploymentCtl) createDeployment(ctx *gin.Context) any {
	deployment := &appsV1.Deployment{}
	query := &requests.CreateDeploymentQuery{}

	athena.Error(ctx.BindJSON(deployment))
	athena.Error(ctx.BindQuery(query))

	if query.Fastmod {
		this.DeploymentService.InitLabel(deployment)
	}

	athena.Error(this.DeploymentService.Create(deployment))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return deployment
}

func (this *DeploymentCtl) updateDeployment(ctx *gin.Context) (v athena.Void) {
	uri := &requests.UpdateDeploymentUri{}
	deployment := &appsV1.Deployment{}

	athena.Error(ctx.BindUri(uri))
	athena.Error(ctx.BindJSON(deployment))
	athena.Error(this.DeploymentService.Update(uri.Namespace, deployment))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
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
	// 获取原生 Deployment
	athena.Handle("GET", "/deployment/:ns/:deployment", this.showDeployment)
	// 查看 Deployment
	athena.Handle("GET", "/deployment/:ns/:deployment/info", this.deploymentInfo)
	// 创建 Deployment
	athena.Handle("POST", "/deployment", this.createDeployment)
	// 编辑 Deployment
	athena.Handle("PUT", "/deployment/:ns", this.updateDeployment)
	// Deployment Pod 列表
	athena.Handle("GET", "/deployment/:ns/:deployment/pods", this.deploymentPods)
	// 删除 Deployment
	athena.Handle("DELETE", "/deployment/:ns/:deployment", this.deleteDeployment)
}
