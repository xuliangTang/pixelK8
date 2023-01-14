package maps

import (
	coreV1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type NodeMap struct {
	data sync.Map // key:nodeName value:*coreV1.Node
}

// Add 添加
func (this *NodeMap) Add(node *coreV1.Node) {
	this.data.Store(node.Name, node)
}

// Update 更新
func (this *NodeMap) Update(node *coreV1.Node) {
	this.data.Store(node.Name, node)
}

// Delete 删除
func (this *NodeMap) Delete(node *coreV1.Node) {
	this.data.Delete(node.Name)
}

// Find 查找node
func (this *NodeMap) Find(nodeName string) *coreV1.Node {
	if node, ok := this.data.Load(nodeName); ok {
		return node.(*coreV1.Node)
	}

	return &coreV1.Node{}
}

// List 获取node列表
func (this *NodeMap) List() []*coreV1.Node {
	ret := make([]*coreV1.Node, 0)
	this.data.Range(func(key, value any) bool {
		ret = append(ret, value.(*coreV1.Node))
		return true
	})

	sort.Sort(coreV1Node(ret))

	return ret
}
