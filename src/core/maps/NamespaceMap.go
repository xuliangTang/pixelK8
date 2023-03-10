package maps

import (
	coreV1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type NamespaceMap struct {
	data sync.Map // key:name value:*coreV1.namespace
}

func (this *NamespaceMap) Add(ns *coreV1.Namespace) {
	this.data.Store(ns.Name, ns)
}

func (this *NamespaceMap) Update(ns *coreV1.Namespace) {
	this.data.Store(ns.Name, ns)
}

func (this *NamespaceMap) Delete(ns *coreV1.Namespace) {
	this.data.Delete(ns.Name)
}

func (this *NamespaceMap) Find(ns string) *coreV1.Namespace {
	if namespace, ok := this.data.Load(ns); ok {
		return namespace.(*coreV1.Namespace)
	}
	return nil
}

func (this *NamespaceMap) List() []*coreV1.Namespace {
	ret := make([]*coreV1.Namespace, 0)
	this.data.Range(func(key, value any) bool {
		ret = append(ret, value.(*coreV1.Namespace))
		return true
	})

	sort.Sort(coreV1Namespace(ret))
	return ret
}
