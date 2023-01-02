package maps

import (
	"crypto/md5"
	"fmt"
	"io"
	coreV1 "k8s.io/api/core/v1"
	"sort"
	"strings"
	"sync"
)

type ConfigmapMap struct {
	data sync.Map
}

// Add 添加
func (this *ConfigmapMap) Add(cm *coreV1.ConfigMap) {
	if cmList, ok := this.data.Load(cm.Namespace); ok {
		cmList = append(cmList.([]*coreV1.ConfigMap), cm)
		this.data.Store(cm.Namespace, cmList)
	} else {
		this.data.Store(cm.Namespace, []*coreV1.ConfigMap{cm})
	}
}

// Update 更新
func (this *ConfigmapMap) Update(cm *coreV1.ConfigMap) bool {
	if cmList, ok := this.data.Load(cm.Namespace); ok {
		cmList := cmList.([]*coreV1.ConfigMap)
		for i, c := range cmList {
			if c.Name == cm.Name && this.Md5Data(c.Data) == this.Md5Data(cm.Data) {
				cmList[i] = cm
				return true
			}
		}
	}

	return false
}

// Delete 删除
func (this *ConfigmapMap) Delete(cm *coreV1.ConfigMap) {
	if cmList, ok := this.data.Load(cm.Namespace); ok {
		cmList := cmList.([]*coreV1.ConfigMap)
		for i, c := range cmList {
			if c.Name == cm.Name {
				newCmList := append(cmList[:i], cmList[i+1:]...)
				this.data.Store(cm.Namespace, newCmList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的configmap列表
func (this *ConfigmapMap) ListByNs(ns string) []*coreV1.ConfigMap {
	if cmList, ok := this.data.Load(ns); ok {
		ret := cmList.([]*coreV1.ConfigMap)
		sort.Sort(coreV1Configmap(ret))
		return ret
	}
	return []*coreV1.ConfigMap{}
}

// Find 查找configmap
func (this *ConfigmapMap) Find(ns, name string) *coreV1.ConfigMap {
	if cmList, ok := this.data.Load(ns); ok {
		for _, c := range cmList.([]*coreV1.ConfigMap) {
			if c.Name == name {
				return c
			}
		}
	}
	return &coreV1.ConfigMap{}
}

// Md5Data 计算map的md5
func (this *ConfigmapMap) Md5Data(data map[string]string) string {
	str := strings.Builder{}
	for k, v := range data {
		str.WriteString(k)
		str.WriteString(v)
	}

	w := md5.New()
	_, _ = io.WriteString(w, str.String())
	return fmt.Sprintf("%x", w.Sum(nil))
}
