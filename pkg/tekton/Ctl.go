package tekton

import (
	"github.com/gin-gonic/gin"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/xuliangTang/athena/athena"
	"net/http"
)

type TektonCtl struct {
	TektonService *TektonService `inject:"-"`
}

func NewTektonCtl() *TektonCtl {
	return &TektonCtl{}
}

func (this *TektonCtl) taskList(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", "default")
	return this.TektonService.LoadTask(ns)
}

func (this *TektonCtl) saveTask(ctx *gin.Context) (v athena.Void) {
	task := &v1beta1.Task{}
	athena.Error(ctx.BindJSON(task))
	athena.Error(this.TektonService.Create(task))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return
}

func (this *TektonCtl) deleteTask(ctx *gin.Context) (v athena.Void) {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.TektonService.Delete(uri.Namespace, uri.Name))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) Build(athena *athena.Athena) {
	// task列表
	athena.Handle(http.MethodGet, "/tekton/tasks", this.taskList)
	// 创建task
	athena.Handle(http.MethodPost, "/tekton/task", this.saveTask)
	// 删除task
	athena.Handle(http.MethodDelete, "/tekton/task/:ns/:name", this.deleteTask)
}
