package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/services"
)

// NamespaceCtl @Controller
type NamespaceCtl struct {
	NsService *services.NamespaceService `inject:"-"`
}

func NewNamespaceCtl() *NamespaceCtl {
	return &NamespaceCtl{}
}

func (this *NamespaceCtl) namespaces(ctx *gin.Context) any {
	nsList := this.NsService.List()
	return nsList
}

func (this *NamespaceCtl) Build(athena *athena.Athena) {
	// 获取ns列表
	athena.Handle("GET", "/namespaces", this.namespaces)
}
