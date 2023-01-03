package services

import (
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// GetLabelsByDepAndRs 通过deployment获取关联ReplicaSet的标签
func (this *CommonService) GetLabelsByDepAndRs(deployment *appsV1.Deployment, rsList []*appsV1.ReplicaSet) ([]map[string]string, error) {
	labels := make([]map[string]string, 0)
	for _, rs := range rsList {
		if this.CheckRsOfDeployment(rs, deployment) {
			selector, err := metaV1.LabelSelectorAsMap(rs.Spec.Selector)
			if err != nil {
				return nil, err
			}
			labels = append(labels, selector)
		}
	}

	return labels, nil
}

// CheckRsOfDeployment 判断rs是否属于当前deployment
func (*CommonService) CheckRsOfDeployment(set *appsV1.ReplicaSet, deployment *appsV1.Deployment) bool {
	for _, rf := range set.OwnerReferences {
		if rf.Kind == "Deployment" && rf.Name == deployment.Name {
			return true
		}
	}

	return false
}

// ParseAnnotations 解析标签
func (*CommonService) ParseAnnotations(annotations string) (ret map[string]string) {
	replace := []string{"\t", " ", "\n", "\r\n"}
	for _, r := range replace {
		annotations = strings.ReplaceAll(annotations, r, "")
	}

	ret = make(map[string]string)
	list := strings.Split(annotations, ";")
	for _, item := range list {
		annotation := strings.Split(item, ":")
		if len(annotation) == 2 {
			ret[annotation[0]] = annotation[1]
		}
	}

	return
}

// SelectorToStrings 转换selector map为切片
func (*CommonService) SelectorToStrings(m map[string]string) (ret []string) {
	ret = make([]string, 0)
	for k, v := range m {
		ret = append(ret, fmt.Sprintf("%s=%s", k, v))
	}

	return ret
}
