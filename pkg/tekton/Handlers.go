package tekton

import (
	"github.com/gin-gonic/gin"
	"log"
	"pixelk8/src/ws"
)

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
