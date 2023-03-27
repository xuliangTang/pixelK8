package tekton

import (
	"github.com/gin-gonic/gin"
	"log"
	"pixelk8/src/ws"
)

// task

type TaskHandler struct {
	TaskMap *TaskMapStruct `inject:"-"`
}

func (this *TaskHandler) OnAdd(obj interface{}) {
	task := ConvertToTask(obj)
	this.TaskMap.Add(task)
	ns := task.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "task",
			"result": gin.H{"ns": ns,
				"data": this.TaskMap.ListAll(ns)},
		},
	)
}

func (this *TaskHandler) OnUpdate(oldObj, newObj interface{}) {
	task := ConvertToTask(newObj)
	err := this.TaskMap.Update(task)
	if err != nil {
		log.Println(err)
		return
	}
	ns := task.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "task",
			"result": gin.H{"ns": ns,
				"data": this.TaskMap.ListAll(ns)},
		},
	)
}

func (this *TaskHandler) OnDelete(obj interface{}) {
	task := ConvertToTask(obj)
	this.TaskMap.Delete(task)
	ns := task.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "task",
			"result": gin.H{"ns": ns,
				"data": this.TaskMap.ListAll(ns)},
		},
	)
}

//Pipeline

type PipelineHandler struct {
	PipelineMap *PipelineMapStruct `inject:"-"`
}

func (this *PipelineHandler) OnAdd(obj interface{}) {
	getObj := ConvertToPipeline(obj)
	this.PipelineMap.Add(getObj)
	ns := getObj.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "pipeline",
			"result": gin.H{"ns": ns,
				"data": this.PipelineMap.ListAll(ns)},
		},
	)
}

func (this *PipelineHandler) OnUpdate(oldObj, newObj interface{}) {
	getObj := ConvertToPipeline(newObj)
	err := this.PipelineMap.Update(getObj)
	if err != nil {
		log.Println(err)
		return
	}
	ns := getObj.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "pipeline",
			"result": gin.H{"ns": ns,
				"data": this.PipelineMap.ListAll(ns)},
		},
	)
}

func (this *PipelineHandler) OnDelete(obj interface{}) {
	getObj := ConvertToPipeline(obj)
	this.PipelineMap.Delete(getObj)
	ns := getObj.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "pipeline",
			"result": gin.H{"ns": ns,
				"data": this.PipelineMap.ListAll(ns)},
		},
	)
}

// PipelineRun

type PipelineRunHandler struct {
	PipelineRunMap *PipelineRunMapStruct `inject:"-"`
}

func (this *PipelineRunHandler) OnAdd(obj interface{}) {
	getObj := ConvertToPipelineRun(obj)
	this.PipelineRunMap.Add(getObj)
	ns := getObj.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "pipelinerun",
			"result": gin.H{"ns": ns,
				"data": this.PipelineRunMap.ListAll(ns)},
		},
	)
}

func (this *PipelineRunHandler) OnUpdate(oldObj, newObj interface{}) {
	getObj := ConvertToPipelineRun(newObj)
	err := this.PipelineRunMap.Update(getObj)
	if err != nil {
		log.Println(err)
		return
	}
	ns := getObj.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "pipelinerun",
			"result": gin.H{"ns": ns,
				"data": this.PipelineRunMap.ListAll(ns)},
		},
	)
}

func (this *PipelineRunHandler) OnDelete(obj interface{}) {
	getObj := ConvertToPipelineRun(obj)
	this.PipelineRunMap.Delete(getObj)
	ns := getObj.Namespace
	ws.ClientMap.SendAll(
		gin.H{
			"type": "pipelinerun",
			"result": gin.H{"ns": ns,
				"data": this.PipelineRunMap.ListAll(ns)},
		},
	)
}
