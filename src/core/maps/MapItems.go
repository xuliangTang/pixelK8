package maps

import "sync"

type MapItems []*MapItem
type MapItem struct {
	Key   string
	Value interface{}
}

func convertToMapItems(m *sync.Map) MapItems {
	items := make(MapItems, 0)
	m.Range(func(key, value interface{}) bool {
		items = append(items, &MapItem{key.(string), value})
		return true
	})
	return items
}

func (this MapItems) Len() int {
	return len(this)
}

func (this MapItems) Less(i, j int) bool {
	return this[i].Key < this[j].Key // 按键排序
}

func (this MapItems) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
