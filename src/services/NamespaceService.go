package services

import (
	"pixelk8/src/core/maps"
	"pixelk8/src/models"
)

// NamespaceService @Service
type NamespaceService struct {
	NsMap *maps.NamespaceMap `inject:"-"`
}

func NewNamespaceService() *NamespaceService {
	return &NamespaceService{}
}

// List 获取ns列表
func (this *NamespaceService) List() (ret []*models.NamespaceModel) {
	nsList := this.NsMap.List()

	for _, ns := range nsList {
		ret = append(ret, &models.NamespaceModel{
			Name: ns.Name,
		})
	}

	return
}
