package maps

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"reflect"
	"sort"
	"sync"
)

type PodMap struct {
	Data sync.Map // key:namespace value:[]*coreV1.Pod
}

// Add 添加
func (this *PodMap) Add(pod *coreV1.Pod) {
	if podList, ok := this.Data.Load(pod.Namespace); ok {
		newList := append(podList.([]*coreV1.Pod), pod)
		this.Data.Store(pod.Namespace, newList)
	} else {
		this.Data.Store(pod.Namespace, []*coreV1.Pod{pod})
	}
}

// Update 更新
func (this *PodMap) Update(pod *coreV1.Pod) error {
	if podList, ok := this.Data.Load(pod.Namespace); ok {
		podList := podList.([]*coreV1.Pod)
		for i, p := range podList {
			if p.Name == pod.Name {
				podList[i] = pod
				break
			}
		}
		return nil
	}

	return fmt.Errorf("replicaset [%s] not found", pod.Name)
}

// Delete 删除
func (this *PodMap) Delete(pod *coreV1.Pod) {
	if podList, ok := this.Data.Load(pod.Namespace); ok {
		podList := podList.([]*coreV1.Pod)
		for i, p := range podList {
			if p.Name == pod.Name {
				newRSList := append(podList[:i], podList[i+1:]...)
				this.Data.Store(pod.Namespace, newRSList)
				break
			}
		}
	}
}

// ListByLabels 根据标签获取Pod列表
func (this *PodMap) ListByLabels(ns string, labels []map[string]string) ([]*coreV1.Pod, error) {
	ret := make([]*coreV1.Pod, 0)
	if podList, ok := this.Data.Load(ns); ok {
		podList := podList.([]*coreV1.Pod)
		for _, p := range podList {
			for _, label := range labels {
				// 判断标签完全匹配
				if reflect.DeepEqual(p.Labels, label) {
					ret = append(ret, p)
				}
			}
		}

		sort.Sort(coreV1Pods(ret))
		return ret, nil
	}

	return []*coreV1.Pod{}, nil
}

// ListByNs 根据ns获取pod列表
func (this *PodMap) ListByNs(ns string) []*coreV1.Pod {
	if podList, ok := this.Data.Load(ns); ok {
		ret := podList.([]*coreV1.Pod)
		sort.Sort(coreV1Pods(ret))

		return ret
	}
	return []*coreV1.Pod{}
}

// Find 查找 Pod
func (this *PodMap) Find(ns string, name string) *coreV1.Pod {
	if podList, ok := this.Data.Load(ns); ok {
		for _, p := range podList.([]*coreV1.Pod) {
			if p.Name == name {
				return p
			}
		}
	}
	return &coreV1.Pod{}
}

// CountByNodeName 根据节点名统计pod数量
func (this *PodMap) CountByNodeName(nodeName string) (count int64) {
	this.Data.Range(func(key, value any) bool {
		pods := value.([]*coreV1.Pod)
		for _, pod := range pods {
			if pod.Spec.NodeName == nodeName {
				count++
			}
		}
		return true
	})

	return
}