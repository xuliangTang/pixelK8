package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"net/http"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
)

// SecretCtl @Controller
type SecretCtl struct {
	SecretService *services.SecretService `inject:"-"`
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}

func (this *SecretCtl) secrets(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	secretList := this.SecretService.ListByNs(ns)

	return this.SecretService.Paging(page, secretList)
}

func (this *SecretCtl) createSecret(ctx *gin.Context) (athena.HttpCode, any) {
	req := &requests.CreateSecret{}
	athena.Error(ctx.BindJSON(req))
	athena.Error(this.SecretService.Create(req))

	return http.StatusCreated, req
}

func (this *SecretCtl) showSecret(ctx *gin.Context) any {
	uri := &requests.ShowSecretUri{}
	athena.Error(ctx.BindUri(uri))

	return this.SecretService.Show(uri)
}

func (this *SecretCtl) deleteSecret(ctx *gin.Context) athena.HttpCode {
	uri := &requests.DeleteSecretUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.SecretService.Delete(uri))

	return http.StatusNoContent
}

func (this *SecretCtl) Build(athena *athena.Athena) {
	// secret列表
	athena.Handle("GET", "/secrets", this.secrets)
	// 创建secret
	athena.Handle("POST", "/secret", this.createSecret)
	// 查看secret
	athena.Handle("GET", "/secret/:ns/:secret", this.showSecret)
	// 删除secret
	athena.Handle("DELETE", "/secret/:ns/:secret", this.deleteSecret)
}
