package maps

import (
	"fmt"
	networkingV1 "k8s.io/api/networking/v1"
	"sort"
	"sync"
)

type IngressMap struct {
	data sync.Map // key:namespace(string), value:[]*networkingV1.Ingress
}

// Add 添加
func (this *IngressMap) Add(ingress *networkingV1.Ingress) {
	if ingList, ok := this.data.Load(ingress.Namespace); ok {
		ingList = append(ingList.([]*networkingV1.Ingress), ingress)
		this.data.Store(ingress.Namespace, ingList)
	} else {
		this.data.Store(ingress.Namespace, []*networkingV1.Ingress{ingress})
	}
}

// Update 更新
func (this *IngressMap) Update(ingress *networkingV1.Ingress) error {
	if ingList, ok := this.data.Load(ingress.Namespace); ok {
		ingList := ingList.([]*networkingV1.Ingress)
		for i, ing := range ingList {
			if ing.Name == ingress.Name {
				ingList[i] = ingress
				break
			}
		}
		return nil
	}

	return fmt.Errorf("ingress [%s] not found", ingress.Name)
}

// Delete 删除
func (this *IngressMap) Delete(ingress *networkingV1.Ingress) {
	if ingList, ok := this.data.Load(ingress.Namespace); ok {
		ingList := ingList.([]*networkingV1.Ingress)
		for i, ing := range ingList {
			if ing.Name == ingress.Name {
				newIngList := append(ingList[:i], ingList[i+1:]...)
				this.data.Store(ingress.Namespace, newIngList)
				break
			}
		}
	}
}

// ListByNs 获取ns下的ingress列表
func (this *IngressMap) ListByNs(ns string) []*networkingV1.Ingress {
	if ingList, ok := this.data.Load(ns); ok {
		ret := ingList.([]*networkingV1.Ingress)
		sort.Sort(networkingV1Ingress(ret))
		return ret
	}
	return []*networkingV1.Ingress{}
}
