package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/services"
)

// DeploymentCtl @controller
type DeploymentCtl struct {
	K8sClient         *kubernetes.Clientset       `inject:"-"`
	DeploymentService *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl) deployments(ctx *gin.Context) any {
	depList := this.DeploymentService.ListByNs("default")
	return depList
}

func (this *DeploymentCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/deployments", this.deployments)
}
