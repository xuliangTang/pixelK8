package tekton

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TektonService struct {
	TaskMap     *TaskMapStruct            `inject:"-"`
	PipelineMap *PipelineMapStruct        `inject:"-"`
	Client      *tektonVersiond.Clientset `inject:"-"`
}

func NewTektonService() *TektonService {
	return &TektonService{}
}

func (this *TektonService) ListTaskByNs(ns string) []*v1beta1.Task {
	return this.TaskMap.ListAll(ns)
}

func (this *TektonService) ShowTask(ns, name string) *v1beta1.Task {
	return this.TaskMap.Find(ns, name)
}

func (this *TektonService) CreateTask(task *v1beta1.Task) error {
	_, err := this.Client.TektonV1beta1().Tasks(task.Namespace).Create(context.Background(), task, v1.CreateOptions{})
	return err
}

func (this *TektonService) UpdateTask(task *v1beta1.Task) error {
	_, err := this.Client.TektonV1beta1().Tasks(task.Namespace).Update(context.Background(), task, v1.UpdateOptions{})
	return err
}

func (this *TektonService) DeleteTask(ns, name string) error {
	return this.Client.TektonV1beta1().Tasks(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}

func (this *TektonService) ListPipelineByNs(ns string) []*v1beta1.Pipeline {
	return this.PipelineMap.ListAll(ns)
}

func (this *TektonService) DeletePipeline(ns, name string) error {
	return this.Client.TektonV1beta1().Pipelines(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}
