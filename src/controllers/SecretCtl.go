package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/properties"
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

func (this *SecretCtl) Build(athena *athena.Athena) {
	// secret列表
	athena.Handle("GET", "/secrets", this.secrets)
}
