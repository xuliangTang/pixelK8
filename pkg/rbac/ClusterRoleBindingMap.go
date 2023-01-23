package rbac

import (
	rbacV1 "k8s.io/api/rbac/v1"
	"sort"
	"sync"
)

type ClusterRoleBindingMap struct {
	data sync.Map // key: name value: *rbacV1.ClusterRoleBinding
}

// Add 添加
func (this *ClusterRoleBindingMap) Add(clusterRoleBinding *rbacV1.ClusterRoleBinding) {
	this.data.Store(clusterRoleBinding.Name, clusterRoleBinding)
}

// Update 更新
func (this *ClusterRoleBindingMap) Update(clusterRoleBinding *rbacV1.ClusterRoleBinding) {
	this.data.Store(clusterRoleBinding.Name, clusterRoleBinding)
}

// Delete 删除
func (this *ClusterRoleBindingMap) Delete(clusterRoleBinding *rbacV1.ClusterRoleBinding) {
	this.data.Delete(clusterRoleBinding.Name)
}

// List 获取clusterRoleBinding列表
func (this *ClusterRoleBindingMap) List() []*rbacV1.ClusterRoleBinding {
	ret := make([]*rbacV1.ClusterRoleBinding, 0)
	this.data.Range(func(key, value any) bool {
		ret = append(ret, value.(*rbacV1.ClusterRoleBinding))
		return true
	})

	sort.Sort(rbacV1ClusterRoleBindings(ret))

	return ret
}

// Find 查找clusterRoleBinding
func (this *ClusterRoleBindingMap) Find(name string) *rbacV1.ClusterRoleBinding {
	if clusterRoleBinding, ok := this.data.Load(name); ok {
		return clusterRoleBinding.(*rbacV1.ClusterRoleBinding)
	}

	return &rbacV1.ClusterRoleBinding{}
}

// 排序clusterRoleBinding 按创建时间
type rbacV1ClusterRoleBindings []*rbacV1.ClusterRoleBinding

func (this rbacV1ClusterRoleBindings) Len() int {
	return len(this)
}

func (this rbacV1ClusterRoleBindings) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this rbacV1ClusterRoleBindings) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

var _ sort.Interface = &rbacV1ClusterRoleBindings{}
