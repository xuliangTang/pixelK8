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
	return this.TektonService.ListTaskByNs(ns)
}

func (this *TektonCtl) showTask(ctx *gin.Context) any {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))

	return this.TektonService.ShowTask(uri.Namespace, uri.Name)
}

func (this *TektonCtl) createTask(ctx *gin.Context) (v athena.Void) {
	task := &v1beta1.Task{}
	athena.Error(ctx.BindJSON(task))
	athena.Error(this.TektonService.CreateTask(task))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return
}

func (this *TektonCtl) updateTask(ctx *gin.Context) (v athena.Void) {
	task := &v1beta1.Task{}
	athena.Error(ctx.BindJSON(task))
	athena.Error(this.TektonService.UpdateTask(task))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) deleteTask(ctx *gin.Context) (v athena.Void) {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.TektonService.DeleteTask(uri.Namespace, uri.Name))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) pipelineList(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", "default")
	return this.TektonService.ListPipelineByNs(ns)
}

func (this *TektonCtl) deletePipeline(ctx *gin.Context) (v athena.Void) {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.TektonService.DeletePipeline(uri.Namespace, uri.Name))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) Build(athena *athena.Athena) {
	// task列表
	athena.Handle(http.MethodGet, "/tekton/tasks", this.taskList)
	// task详情
	athena.Handle(http.MethodGet, "/tekton/task/:ns/:name", this.showTask)
	// 创建task
	athena.Handle(http.MethodPost, "/tekton/task", this.createTask)
	// 更新task
	athena.Handle(http.MethodPut, "/tekton/task", this.updateTask)
	// 删除task
	athena.Handle(http.MethodDelete, "/tekton/task/:ns/:name", this.deleteTask)

	// pipeline列表
	athena.Handle(http.MethodGet, "/tekton/pipelines", this.pipelineList)
	// 删除pipeline
	athena.Handle(http.MethodDelete, "/tekton/pipeline/:ns/:name", this.deletePipeline)
}
