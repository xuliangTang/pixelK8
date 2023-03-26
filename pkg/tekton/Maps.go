package tekton

import (
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"sort"
	"sync"
)

type V1Task []*v1beta1.Task

func (this V1Task) Len() int {
	return len(this)
}
func (this V1Task) Less(i, j int) bool {
	// 根据时间排序倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this V1Task) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type TaskMapStruct struct {
	data sync.Map // [ns string] []*v1beta1.Task
}

func (this *TaskMapStruct) Add(item *v1beta1.Task) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*v1beta1.Task), item)
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*v1beta1.Task{item})
	}
}

func (this *TaskMapStruct) Update(item *v1beta1.Task) error {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, rangeItem := range list.([]*v1beta1.Task) {
			if rangeItem.Name == item.Name {
				list.([]*v1beta1.Task)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("task-%s not found", item.Name)
}

func (this *TaskMapStruct) Delete(svc *v1beta1.Task) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, rangeItem := range list.([]*v1beta1.Task) {
			if rangeItem.Name == svc.Name {
				newList := append(list.([]*v1beta1.Task)[:i], list.([]*v1beta1.Task)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}

func (this *TaskMapStruct) ListAll(ns string) []*v1beta1.Task {
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*v1beta1.Task)
		sort.Sort(V1Task(newList)) // 按时间倒排序
		return newList
	}
	return []*v1beta1.Task{} //返回空列表
}
