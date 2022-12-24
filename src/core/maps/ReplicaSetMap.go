package maps

import (
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	"sync"
)

type ReplicaSetMap struct {
	Data sync.Map // key:namespace value:[]*appsV1.ReplicaSet
}

// Add 添加
func (this *ReplicaSetMap) Add(set *appsV1.ReplicaSet) {
	if rsList, ok := this.Data.Load(set.Namespace); ok {
		newList := append(rsList.([]*appsV1.ReplicaSet), set)
		this.Data.Store(set.Namespace, newList)
	} else {
		this.Data.Store(set.Namespace, []*appsV1.ReplicaSet{set})
	}
}

// Update 更新
func (this *ReplicaSetMap) Update(set *appsV1.ReplicaSet) error {
	if rsList, ok := this.Data.Load(set.Namespace); ok {
		rsList := rsList.([]*appsV1.ReplicaSet)
		for i, rs := range rsList {
			if rs.Name == set.Name {
				rsList[i] = set
				break
			}
		}
		return nil
	}

	return fmt.Errorf("replicaset [%s] not found", set.Name)
}

// Delete 删除
func (this *ReplicaSetMap) Delete(set *appsV1.ReplicaSet) {
	if rsList, ok := this.Data.Load(set.Namespace); ok {
		rsList := rsList.([]*appsV1.ReplicaSet)
		for i, rs := range rsList {
			if rs.Name == set.Name {
				newRSList := append(rsList[:i], rsList[i+1:]...)
				this.Data.Store(set.Namespace, newRSList)
				break
			}
		}
	}
}

// ListByNs 获取列表
func (this *ReplicaSetMap) ListByNs(ns string) ([]*appsV1.ReplicaSet, error) {
	if rsList, ok := this.Data.Load(ns); ok {
		return rsList.([]*appsV1.ReplicaSet), nil
	}

	return []*appsV1.ReplicaSet{}, nil
}
