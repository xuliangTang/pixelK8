package rbac

import (
	rbacV1 "k8s.io/api/rbac/v1"
	"sort"
	"sync"
)

type ClusterRoleMap struct {
	data sync.Map // key: name value: *rbacV1.ClusterRole
}

// Add 添加
func (this *ClusterRoleMap) Add(clusterRole *rbacV1.ClusterRole) {
	this.data.Store(clusterRole.Name, clusterRole)
}

// Update 更新
func (this *ClusterRoleMap) Update(clusterRole *rbacV1.ClusterRole) {
	this.data.Store(clusterRole.Name, clusterRole)
}

// Delete 删除
func (this *ClusterRoleMap) Delete(clusterRole *rbacV1.ClusterRole) {
	this.data.Delete(clusterRole.Name)
}

// List 获取clusterRole列表
func (this *ClusterRoleMap) List() []*rbacV1.ClusterRole {
	ret := make([]*rbacV1.ClusterRole, 0)
	this.data.Range(func(key, value any) bool {
		ret = append(ret, value.(*rbacV1.ClusterRole))
		return true
	})

	sort.Sort(rbacV1ClusterRoles(ret))

	return ret
}

// Find 查找clusterRole
func (this *ClusterRoleMap) Find(name string) *rbacV1.ClusterRole {
	if clusterRole, ok := this.data.Load(name); ok {
		return clusterRole.(*rbacV1.ClusterRole)
	}

	return &rbacV1.ClusterRole{}
}

// 排序clusterRole 按创建时间
type rbacV1ClusterRoles []*rbacV1.ClusterRole

func (this rbacV1ClusterRoles) Len() int {
	return len(this)
}

func (this rbacV1ClusterRoles) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this rbacV1ClusterRoles) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
