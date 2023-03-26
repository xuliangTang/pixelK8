package tekton

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
)

type TektonCtl struct {
	TektonService *TektonService `inject:"-"`
}

func NewTektonCtl() *TektonCtl {
	return &TektonCtl{}
}

func (this *TektonCtl) TaskList(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", "default")
	return this.TektonService.LoadTask(ns)
}

func (this *TektonCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/tekton/tasks", this.TaskList)
}
