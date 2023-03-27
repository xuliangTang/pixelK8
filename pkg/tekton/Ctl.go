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

func (this *TektonCtl) showPipeline(ctx *gin.Context) any {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))

	return this.TektonService.ShowPipeline(uri.Namespace, uri.Name)
}

func (this *TektonCtl) createPipeline(ctx *gin.Context) (v athena.Void) {
	pipeline := &v1beta1.Pipeline{}
	athena.Error(ctx.BindJSON(pipeline))
	athena.Error(this.TektonService.CreatePipeline(pipeline))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) updatePipeline(ctx *gin.Context) (v athena.Void) {
	pipeline := &v1beta1.Pipeline{}
	athena.Error(ctx.BindJSON(pipeline))
	athena.Error(this.TektonService.UpdatePipeline(pipeline))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return
}

func (this *TektonCtl) deletePipeline(ctx *gin.Context) (v athena.Void) {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.TektonService.DeletePipeline(uri.Namespace, uri.Name))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) pipelineRunList(ctx *gin.Context) any {
	ns := ctx.DefaultQuery("ns", "default")
	return this.TektonService.ListPipelineRunByNs(ns)
}

func (this *TektonCtl) showPipelineRun(ctx *gin.Context) any {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))

	return this.TektonService.ShowPipelineRun(uri.Namespace, uri.Name)
}

func (this *TektonCtl) createPipelineRun(ctx *gin.Context) (v athena.Void) {
	pipelineRun := &v1beta1.PipelineRun{}
	athena.Error(ctx.BindJSON(pipelineRun))
	athena.Error(this.TektonService.CreatePipelineRun(pipelineRun))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusNoContent)
	return
}

func (this *TektonCtl) updatePipelineRun(ctx *gin.Context) (v athena.Void) {
	pipelineRun := &v1beta1.PipelineRun{}
	athena.Error(ctx.BindJSON(pipelineRun))
	athena.Error(this.TektonService.UpdatePipelineRun(pipelineRun))

	ctx.Set(athena.CtxHttpStatusCode, http.StatusCreated)
	return
}

func (this *TektonCtl) deletePipelineRun(ctx *gin.Context) (v athena.Void) {
	uri := &RsUri{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(this.TektonService.DeletePipelineRun(uri.Namespace, uri.Name))

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
	// pipeline详情
	athena.Handle(http.MethodGet, "/tekton/pipeline/:ns/:name", this.showPipeline)
	// 创建pipeline
	athena.Handle(http.MethodPost, "/tekton/pipeline", this.createPipeline)
	// 更新pipeline
	athena.Handle(http.MethodPut, "/tekton/pipeline", this.updatePipeline)
	// 删除pipeline
	athena.Handle(http.MethodDelete, "/tekton/pipeline/:ns/:name", this.deletePipeline)

	// pipelineRun列表
	athena.Handle(http.MethodGet, "/tekton/pipelineruns", this.pipelineRunList)
	// pipelineRun详情
	athena.Handle(http.MethodGet, "/tekton/pipelinerun/:ns/:name", this.showPipelineRun)
	// 创建pipelineRun
	athena.Handle(http.MethodPost, "/tekton/pipelinerun", this.createPipelineRun)
	// 更新pipelineRun
	athena.Handle(http.MethodPut, "/tekton/pipelinerun", this.updatePipelineRun)
	// 删除pipelineRun
	athena.Handle(http.MethodDelete, "/tekton/pipelinerun/:ns/:name", this.deletePipelineRun)
}
