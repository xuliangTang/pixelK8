package rbac

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type ServiceAccountMap struct {
	data sync.Map // key: namespace value:coreV1.ServiceAccount
}

// Add 添加
func (this *ServiceAccountMap) Add(serviceAccount *coreV1.ServiceAccount) {
	if saList, ok := this.data.Load(serviceAccount.Namespace); ok {
		saList = append(saList.([]*coreV1.ServiceAccount), serviceAccount)
		this.data.Store(serviceAccount.Namespace, saList)
	} else {
		this.data.Store(serviceAccount.Namespace, []*coreV1.ServiceAccount{serviceAccount})
	}
}

// Update 更新
func (this *ServiceAccountMap) Update(serviceAccount *coreV1.ServiceAccount) error {
	if saList, ok := this.data.Load(serviceAccount.Namespace); ok {
		saList := saList.([]*coreV1.ServiceAccount)
		for i, sa := range saList {
			if sa.Name == serviceAccount.Name {
				saList[i] = serviceAccount
				break
			}
		}
		return nil
	}

	return fmt.Errorf("serviceAccount [%s] not found", serviceAccount.Name)
}

// Delete 删除
func (this *ServiceAccountMap) Delete(serviceAccount *coreV1.ServiceAccount) {
	if saList, ok := this.data.Load(serviceAccount.Namespace); ok {
		saList := saList.([]*coreV1.ServiceAccount)
		for i, sa := range saList {
			if sa.Name == serviceAccount.Name {
				newSaList := append(saList[:i], saList[i+1:]...)
				this.data.Store(serviceAccount.Namespace, newSaList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的serviceAccount列表
func (this *ServiceAccountMap) ListByNs(ns string) []*coreV1.ServiceAccount {
	if saList, ok := this.data.Load(ns); ok {
		ret := saList.([]*coreV1.ServiceAccount)
		sort.Sort(coreV1ServiceAccounts(ret))
		return ret
	}

	return []*coreV1.ServiceAccount{}
}

// Find 查找
func (this *ServiceAccountMap) Find(ns, saName string) *coreV1.ServiceAccount {
	if saList, ok := this.data.Load(ns); ok {
		for _, sa := range saList.([]*coreV1.ServiceAccount) {
			if sa.Name == saName {
				return sa
			}
		}
	}

	return &coreV1.ServiceAccount{}
}

// 排序serviceAccount 按创建时间
type coreV1ServiceAccounts []*coreV1.ServiceAccount

func (this coreV1ServiceAccounts) Len() int {
	return len(this)
}

func (this coreV1ServiceAccounts) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this coreV1ServiceAccounts) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
