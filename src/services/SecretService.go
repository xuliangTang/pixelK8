package services

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
)

// SecretService @Service
type SecretService struct {
	SecretMap *maps.SecretMap `inject:"-"`
	Localize  *i18n.Localizer `inject:"-"`
}

func NewSecretService() *SecretService {
	return &SecretService{}
}

// ListByNs 获取ns下的secret列表
func (this *SecretService) ListByNs(ns string) (ret []*dto.SecretList) {
	secretList := this.SecretMap.ListByNs(ns)

	ret = make([]*dto.SecretList, len(secretList))
	for i, secret := range secretList {
		typeStr := athena.UnwrapOrEmpty(this.Localize.Localize(&i18n.LocalizeConfig{
			MessageID: string("k8s.secret.type." + secret.Type),
		}))
		ret[i] = &dto.SecretList{
			Name:      secret.Name,
			Namespace: secret.Namespace,
			Type:      [2]string{string(secret.Type), typeStr},
			CreatedAt: secret.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页secret切片
func (this *SecretService) Paging(page *athena.Page, secretList []*dto.SecretList) athena.Collection {
	var count int
	iSecList := make([]any, len(secretList))
	for i, sec := range iSecList {
		count++
		iSecList[i] = sec
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iSecList)
	collection := athena.NewCollection(secretList[start:end], page)
	return *collection
}
