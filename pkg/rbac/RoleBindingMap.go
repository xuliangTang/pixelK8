package rbac

import (
	"fmt"
	rbacV1 "k8s.io/api/rbac/v1"
	"sort"
	"sync"
)

type RoleBindingMap struct {
	data sync.Map // key: namespace value: []*rbacV1.RoleBinding
}

// Add 添加
func (this *RoleBindingMap) Add(roleBinding *rbacV1.RoleBinding) {
	if roleBindingList, ok := this.data.Load(roleBinding.Namespace); ok {
		roleBindingList = append(roleBindingList.([]*rbacV1.RoleBinding), roleBinding)
		this.data.Store(roleBinding.Namespace, roleBindingList)
	} else {
		this.data.Store(roleBinding.Namespace, []*rbacV1.RoleBinding{roleBinding})
	}
}

// Update 更新
func (this *RoleBindingMap) Update(roleBinding *rbacV1.RoleBinding) error {
	if roleBindingList, ok := this.data.Load(roleBinding.Namespace); ok {
		roleBindingList := roleBindingList.([]*rbacV1.RoleBinding)
		for i, rb := range roleBindingList {
			if rb.Name == roleBinding.Name {
				roleBindingList[i] = roleBinding
				break
			}
		}
		return nil
	}

	return fmt.Errorf("roleBinding [%s] not found", roleBinding.Name)
}

// Delete 删除
func (this *RoleBindingMap) Delete(roleBinding *rbacV1.RoleBinding) {
	if roleBindingList, ok := this.data.Load(roleBinding.Namespace); ok {
		roleBindingList := roleBindingList.([]*rbacV1.RoleBinding)
		for i, rb := range roleBindingList {
			if rb.Name == roleBinding.Name {
				newRoleBindingList := append(roleBindingList[:i], roleBindingList[i+1:]...)
				this.data.Store(roleBinding.Namespace, newRoleBindingList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的roleBinding列表
func (this *RoleBindingMap) ListByNs(ns string) []*rbacV1.RoleBinding {
	if roleBindingList, ok := this.data.Load(ns); ok {
		ret := roleBindingList.([]*rbacV1.RoleBinding)
		sort.Sort(rbacV1RoleBindings(ret))
		return ret
	}

	return []*rbacV1.RoleBinding{}
}

// Find 查找roleBinding
func (this *RoleBindingMap) Find(ns, roleBindingName string) *rbacV1.RoleBinding {
	if roleBindingList, ok := this.data.Load(ns); ok {
		for _, roleBinding := range roleBindingList.([]*rbacV1.RoleBinding) {
			if roleBinding.Name == roleBindingName {
				return roleBinding
			}
		}
	}

	return &rbacV1.RoleBinding{}
}

// 排序roleBinding 按创建时间
type rbacV1RoleBindings []*rbacV1.RoleBinding

func (this rbacV1RoleBindings) Len() int {
	return len(this)
}

func (this rbacV1RoleBindings) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this rbacV1RoleBindings) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
