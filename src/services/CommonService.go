package services

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"strings"
)

// CommonService @service
type CommonService struct{}

func NewCommonService() *CommonService {
	return &CommonService{}
}

// GetImageShortName 获取简短镜像名
func (*CommonService) GetImageShortName(cs []coreV1.Container) string {
	var images strings.Builder
	images.WriteString(cs[0].Image)
	if lenImage := len(cs); lenImage > 1 {
		images.WriteString(fmt.Sprintf("和%d个镜像", lenImage))
	}

	return images.String()
}
