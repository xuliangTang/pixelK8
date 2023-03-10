package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/requests"
)

// DeploymentService @service
type DeploymentService struct {
	DeploymentMap *maps.DeploymentMap   `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	K8sClient     *kubernetes.Clientset `inject:"-"`
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
			CreatedAt:   dep.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}

	return
}

// Paging 分页deployments切片
func (this *DeploymentService) Paging(page *athena.Page, depList []*dto.DeploymentList) athena.Collection {
	var count, countReady int
	iDepList := make([]any, len(depList))
	for i, dep := range depList {
		iDepList[i] = dep
		count++
		if dep.IsCompleted {
			countReady++
		}
	}

	page.Extend = gin.H{"count": count, "count_ready": countReady}
	// 分页
	start, end := page.SlicePage(iDepList)
	collection := athena.NewCollection(depList[start:end], page)
	return *collection
}

// Show 获取deployment原生对象
func (this *DeploymentService) Show(uri *requests.ShowDeploymentUri) *appsV1.Deployment {
	dep, err := this.DeploymentMap.Find(uri.Namespace, uri.Deployment)
	if err != nil {
		return nil
	}

	return dep
}

// Info 查看deployment
func (this *DeploymentService) Info(uri *requests.ShowDeploymentUri) dto.DeploymentInfo {
	dep, err := this.DeploymentMap.Find(uri.Namespace, uri.Deployment)
	if err != nil {
		return dto.DeploymentInfo{}
	}

	return dto.DeploymentInfo{
		Name:        dep.Name,
		Namespace:   dep.Namespace,
		Replicas:    [3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas},
		Images:      this.CommonService.GetImageShortName(dep.Spec.Template.Spec.Containers),
		IsCompleted: this.checkDeploymentIsCompleted(dep),
		Message:     this.getDeploymentConditionsMessage(dep),
		CreatedAt:   dep.CreationTimestamp.Format("2006-01-02 15:04:05"),
	}
}

// Create 创建deployment
func (this *DeploymentService) Create(deploy *appsV1.Deployment) error {
	_, err := this.K8sClient.AppsV1().Deployments(deploy.Namespace).
		Create(context.Background(), deploy, metaV1.CreateOptions{})
	return err
}

// InitLabel 初始化deployment的matchLabel
func (this *DeploymentService) InitLabel(deploy *appsV1.Deployment) {
	if deploy.Spec.Selector == nil {
		deploy.Spec.Selector = &metaV1.LabelSelector{MatchLabels: map[string]string{"app": deploy.Name}}
	}
	if deploy.Spec.Selector.MatchLabels == nil {
		deploy.Spec.Selector.MatchLabels = map[string]string{"app": deploy.Name}
	}
	if deploy.Spec.Template.ObjectMeta.Labels == nil {
		deploy.Spec.Template.ObjectMeta.Labels = map[string]string{"app": deploy.Name}
	}
	deploy.Spec.Selector.MatchLabels["app"] = deploy.Name

	deploy.Spec.Template.ObjectMeta.Labels["app"] = deploy.Name
}

// Update 编辑deployment
func (this *DeploymentService) Update(ns string, deploy *appsV1.Deployment) error {
	_, err := this.K8sClient.AppsV1().Deployments(ns).
		Update(context.Background(), deploy, metaV1.UpdateOptions{})
	return err
}

// Delete 删除deployment
func (this *DeploymentService) Delete(uri *requests.DeleteDeploymentUri) error {
	return this.K8sClient.AppsV1().Deployments(uri.Namespace).
		Delete(context.Background(), uri.Name, metaV1.DeleteOptions{})
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
