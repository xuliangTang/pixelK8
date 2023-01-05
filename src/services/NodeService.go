package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
)

// NodeService @Service
type NodeService struct {
	NodeMap *maps.NodeMap `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

// List node列表
func (this *NodeService) List() (ret []*dto.NodeList) {
	nodeList := this.NodeMap.List()

	ret = make([]*dto.NodeList, len(nodeList))
	for i, node := range nodeList {
		ret[i] = &dto.NodeList{
			Name:      node.Name,
			Ip:        node.Status.Addresses[0].Address,
			CreatedAt: node.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页node切片
func (this *NodeService) Paging(page *athena.Page, nodeList []*dto.NodeList) athena.Collection {
	var count int
	iNodeList := make([]any, len(nodeList))
	for i, node := range iNodeList {
		count++
		iNodeList[i] = node
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iNodeList)
	collection := athena.NewCollection(nodeList[start:end], page)
	return *collection
}
