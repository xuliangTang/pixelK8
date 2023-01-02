package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/requests"
)

// PodService @Service
type PodService struct {
	DeploymentMap *maps.DeploymentMap `inject:"-"`
	ReplicaSetMap *maps.ReplicaSetMap `inject:"-"`
	PodMap        *maps.PodMap        `inject:"-"`
	EventMap      *maps.EventMap      `inject:"-"`
	CommonService *CommonService      `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

// ListByDeployment 获取deployment的pod列表
func (this *PodService) ListByDeployment(ns, depName string) []*dto.PodList {
	deployment := athena.Unwrap(this.DeploymentMap.Find(ns, depName)).(*appsV1.Deployment)
	rsList := athena.Unwrap(this.ReplicaSetMap.ListByNs(ns)).([]*appsV1.ReplicaSet)
	labels := athena.Unwrap(this.CommonService.GetLabelsByDepAndRs(deployment, rsList)).([]map[string]string)
	return this.ListByLabels(ns, labels)
}

// ListByLabels 通过labels获取pods列表
func (this *PodService) ListByLabels(ns string, labels []map[string]string) (ret []*dto.PodList) {
	podList := athena.Unwrap(this.PodMap.ListByLabels(ns, labels)).([]*coreV1.Pod)

	return this.convertPodList(podList)
}

// ListByNs 通过ns获取pod列表
func (this *PodService) ListByNs(ns string) (ret []*dto.PodList) {
	podList := this.PodMap.ListByNs(ns)

	return this.convertPodList(podList)
}

// Paging 将pods切片数据分页
func (this *PodService) Paging(page *athena.Page, podList []*dto.PodList) athena.Collection {
	var count, countReady int // pod总数和就绪数
	iPodList := make([]any, len(podList))
	for i, pod := range podList {
		iPodList[i] = pod
		count++
		if pod.IsReady {
			countReady++
		}
	}

	page.Extend = gin.H{"count": count, "count_ready": countReady}
	// 分页
	start, end := page.SlicePage(iPodList)
	collection := athena.NewCollection(podList[start:end], page)
	return *collection
}

// CheckPodReady 评估Pod是否就绪
func (*PodService) CheckPodReady(pod *coreV1.Pod) bool {
	if pod.Status.Phase != "Running" {
		return false
	}

	// 所有容器是否就绪
	for _, condition := range pod.Status.Conditions {
		if condition.Status != "True" {
			return false
		}
	}

	// readinessGates
	for _, rg := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == rg.ConditionType && condition.Status != "True" {
				return false
			}
		}
	}

	return true
}

// GetPodContainers 获取pod所有容器
func (this *PodService) GetPodContainers(uri *requests.PodAllContainersUri) (ret []*dto.PodContainerList) {
	pod := this.PodMap.Find(uri.Namespace, uri.Name)
	ret = make([]*dto.PodContainerList, 0)
	for _, c := range pod.Spec.Containers {
		ret = append(ret, &dto.PodContainerList{
			Name: c.Name,
		})
	}

	return
}

// 将原生pods列表转换为dto对象
func (this *PodService) convertPodList(podList []*coreV1.Pod) (ret []*dto.PodList) {
	ret = make([]*dto.PodList, len(podList))
	for i, pod := range podList {
		ret[i] = &dto.PodList{
			Name:      pod.Name,
			NameSpace: pod.Namespace,
			NodeName:  pod.Spec.NodeName,
			Images:    this.CommonService.GetImageShortName(pod.Spec.Containers),
			Ip:        [2]string{pod.Status.PodIP, pod.Status.HostIP},
			Phase:     string(pod.Status.Phase),
			IsReady:   this.CheckPodReady(pod),
			Message:   this.EventMap.GetMessage(pod.Namespace, "Pod", pod.Name),
			CreatedAt: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}

	return
}
