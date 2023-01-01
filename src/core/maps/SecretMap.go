package maps

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type SecretMap struct {
	data sync.Map
}

// Add 添加
func (this *SecretMap) Add(secret *coreV1.Secret) {
	if secretList, ok := this.data.Load(secret.Namespace); ok {
		secretList = append(secretList.([]*coreV1.Secret), secret)
		this.data.Store(secret.Namespace, secretList)
	} else {
		this.data.Store(secret.Namespace, []*coreV1.Secret{secret})
	}
}

// Update 更新
func (this *SecretMap) Update(secret *coreV1.Secret) error {
	if secretList, ok := this.data.Load(secret.Namespace); ok {
		secretList := secretList.([]*coreV1.Secret)
		for i, sec := range secretList {
			if sec.Name == secret.Name {
				secretList[i] = secret
				break
			}
		}
		return nil
	}

	return fmt.Errorf("secret [%s] not found", secret.Name)
}

// Delete 删除
func (this *SecretMap) Delete(secret *coreV1.Secret) {
	if secretList, ok := this.data.Load(secret.Namespace); ok {
		secretList := secretList.([]*coreV1.Secret)
		for i, sec := range secretList {
			if sec.Name == secret.Name {
				newSecList := append(secretList[:i], secretList[i+1:]...)
				this.data.Store(secret.Namespace, newSecList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的secret列表
func (this *SecretMap) ListByNs(ns string) []*coreV1.Secret {
	if secList, ok := this.data.Load(ns); ok {
		ret := secList.([]*coreV1.Secret)
		sort.Sort(coreV1Secret(ret))
		return ret
	}
	return []*coreV1.Secret{}
}

// Find 查找secret
func (this *SecretMap) Find(ns, name string) *coreV1.Secret {
	if secList, ok := this.data.Load(ns); ok {
		for _, sec := range secList.([]*coreV1.Secret) {
			if sec.Name == name {
				return sec
			}
		}
	}
	return &coreV1.Secret{}
}
