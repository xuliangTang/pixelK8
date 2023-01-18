package rbac

import (
	"fmt"
	rbacV1 "k8s.io/api/rbac/v1"
	"sort"
	"sync"
)

type RoleMap struct {
	data sync.Map // key: namespace value: rbacV1.Role
}

// Add 添加
func (this *RoleMap) Add(role *rbacV1.Role) {
	if roleList, ok := this.data.Load(role.Namespace); ok {
		roleList = append(roleList.([]*rbacV1.Role), role)
		this.data.Store(role.Namespace, roleList)
	} else {
		this.data.Store(role.Namespace, []*rbacV1.Role{role})
	}
}

// Update 更新
func (this *RoleMap) Update(role *rbacV1.Role) error {
	if roleList, ok := this.data.Load(role.Namespace); ok {
		roleList := roleList.([]*rbacV1.Role)
		for i, r := range roleList {
			if r.Name == role.Name {
				roleList[i] = role
				break
			}
		}
		return nil
	}

	return fmt.Errorf("role [%s] not found", role.Name)
}

// Delete 删除
func (this *RoleMap) Delete(role *rbacV1.Role) {
	if roleList, ok := this.data.Load(role.Namespace); ok {
		roleList := roleList.([]*rbacV1.Role)
		for i, r := range roleList {
			if r.Name == role.Name {
				newRoleList := append(roleList[:i], roleList[i+1:]...)
				this.data.Store(role.Namespace, newRoleList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的role列表
func (this *RoleMap) ListByNs(ns string) []*rbacV1.Role {
	if roleList, ok := this.data.Load(ns); ok {
		ret := roleList.([]*rbacV1.Role)
		sort.Sort(rbacV1Roles(ret))
		return ret
	}

	return []*rbacV1.Role{}
}

// Find 查找
func (this *RoleMap) Find(ns, roleName string) *rbacV1.Role {
	if roleList, ok := this.data.Load(ns); ok {
		for _, role := range roleList.([]*rbacV1.Role) {
			if role.Name == roleName {
				return role
			}
		}
	}

	return &rbacV1.Role{}
}

// 排序role 按创建时间
type rbacV1Roles []*rbacV1.Role

func (this rbacV1Roles) Len() int {
	return len(this)
}

func (this rbacV1Roles) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this rbacV1Roles) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
