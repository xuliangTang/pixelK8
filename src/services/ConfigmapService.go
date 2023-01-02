package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/requests"
)

// ConfigmapService @Service
type ConfigmapService struct {
	CmMap     *maps.ConfigmapMap    `inject:"-"`
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewConfigmapService() *ConfigmapService {
	return &ConfigmapService{}
}

// ListByNs 获取ns下的configmaps
func (this *ConfigmapService) ListByNs(ns string) (ret []*dto.ConfigmapList) {
	cmList := this.CmMap.ListByNs(ns)

	ret = make([]*dto.ConfigmapList, len(cmList))
	for i, cm := range cmList {
		ret[i] = &dto.ConfigmapList{
			Name:      cm.Name,
			Namespace: cm.Namespace,
			CreatedAt: cm.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页configmap切片
func (this *ConfigmapService) Paging(page *athena.Page, cmList []*dto.ConfigmapList) athena.Collection {
	var count int
	iCmList := make([]any, len(cmList))
	for i, cm := range iCmList {
		count++
		iCmList[i] = cm
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iCmList)
	collection := athena.NewCollection(cmList[start:end], page)
	return *collection
}

// Create 创建configmap
func (this *ConfigmapService) Create(req *requests.CreateConfigmap) error {
	configmap := &coreV1.ConfigMap{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
		},
		Data: req.Data,
	}

	_, err := this.K8sClient.CoreV1().ConfigMaps(req.Namespace).
		Create(context.Background(), configmap, metaV1.CreateOptions{})

	return err
}
