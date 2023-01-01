package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xuliangTang/athena/athena"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/helpers"
	"pixelk8/src/requests"
)

// SecretService @Service
type SecretService struct {
	SecretMap *maps.SecretMap       `inject:"-"`
	Localize  *i18n.Localizer       `inject:"-"`
	K8sClient *kubernetes.Clientset `inject:"-"`
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

// Create 创建secret
func (this *SecretService) Create(req *requests.CreateSecret) error {
	secret := &coreV1.Secret{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
		},
		Type:       coreV1.SecretType(req.Type),
		StringData: req.Data,
	}

	_, err := this.K8sClient.CoreV1().Secrets(req.Namespace).
		Create(context.Background(), secret, metaV1.CreateOptions{})

	return err
}

// Show 查看secret
func (this *SecretService) Show(uri *requests.ShowSecretUri) *dto.SecretShow {
	secret := this.SecretMap.Find(uri.Namespace, uri.Name)

	typeStr := athena.UnwrapOrEmpty(this.Localize.Localize(&i18n.LocalizeConfig{
		MessageID: string("k8s.secret.type." + secret.Type),
	}))

	return &dto.SecretShow{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		Type:      [2]string{string(secret.Type), typeStr},
		Data:      secret.Data,
		CreatedAt: secret.CreationTimestamp.Format(athena.DateTimeFormat),
		Ext:       this.ParseExt(secret),
	}
}

// ParseExt 解析secret扩展信息
func (this *SecretService) ParseExt(secret *coreV1.Secret) any {
	if string(secret.Type) == string(coreV1.SecretTypeTLS) {
		if crt, ok := secret.Data["tls.crt"]; ok {
			crtModel := helpers.ParseCert(crt)
			if crtModel != nil {
				return crtModel
			}
		}
	}

	return struct{}{}
}
