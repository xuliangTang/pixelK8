package services

import (
	"github.com/xuliangTang/athena/athena"
	appsV1 "k8s.io/api/apps/v1"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
)

// DeploymentService @service
type DeploymentService struct {
	DeploymentMap *maps.DeploymentMap `inject:"-"`
	CommonService *CommonService      `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}

// ListByNs 获取deployment列表
func (this *DeploymentService) ListByNs(namespace string) (ret []*dto.DeploymentList) {
	depList := athena.Unwrap(this.DeploymentMap.ListByNs(namespace)).([]*appsV1.Deployment)

	for _, dep := range depList {
		ret = append(ret, &dto.DeploymentList{
			Name:        dep.Name,
			Namespace:   dep.Namespace,
			Replicas:    [3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas},
			Images:      this.CommonService.GetImageShortName(dep.Spec.Template.Spec.Containers),
			IsCompleted: this.checkDeploymentIsCompleted(dep),
			Message:     this.getDeploymentConditionsMessage(dep),
		})
	}

	return
}

// 评估deployment是否就绪
func (*DeploymentService) checkDeploymentIsCompleted(deployment *appsV1.Deployment) bool {
	return deployment.Status.Replicas == deployment.Status.AvailableReplicas
}

// 从Status.Conditions中获取deployment失败信息
func (*DeploymentService) getDeploymentConditionsMessage(deployment *appsV1.Deployment) string {
	for _, condition := range deployment.Status.Conditions {
		if condition.Type == "Available" && condition.Status != "True" {
			return condition.Message
		}
	}

	return ""
}
