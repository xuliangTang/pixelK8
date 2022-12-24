package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
)

// DeploymentCtl @controller
type DeploymentCtl struct {
	K8sClient         *kubernetes.Clientset       `inject:"-"`
	DeploymentService *services.DeploymentService `inject:"-"`
	PodService        *services.PodService        `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl) deployments(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	depList := this.DeploymentService.ListByNs(ns)
	return depList
}

func (this *DeploymentCtl) deploymentPods(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	var uri requests.DeploymentUri
	athena.Error(ctx.BindUri(&uri))
	podList := this.PodService.ListByDeployment(ns, uri.Deployment)
	return podList
}

func (this *DeploymentCtl) Build(athena *athena.Athena) {
	// Deployment 列表
	athena.Handle("GET", "/deployments", this.deployments)
	// Deployment Pod 列表
	athena.Handle("GET", "/deployment/:deployment/pods", this.deploymentPods)
}
