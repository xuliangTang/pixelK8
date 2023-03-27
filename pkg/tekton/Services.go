package tekton

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TektonService struct {
	TaskMap        *TaskMapStruct            `inject:"-"`
	PipelineMap    *PipelineMapStruct        `inject:"-"`
	PipelineRunMap *PipelineRunMapStruct     `inject:"-"`
	Client         *tektonVersiond.Clientset `inject:"-"`
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

func (this *TektonService) ShowPipeline(ns, name string) *v1beta1.Pipeline {
	return this.PipelineMap.Find(ns, name)
}

func (this *TektonService) CreatePipeline(pipeline *v1beta1.Pipeline) error {
	_, err := this.Client.TektonV1beta1().Pipelines(pipeline.Namespace).Create(context.Background(), pipeline, v1.CreateOptions{})
	return err
}

func (this *TektonService) UpdatePipeline(pipeline *v1beta1.Pipeline) error {
	_, err := this.Client.TektonV1beta1().Pipelines(pipeline.Namespace).Update(context.Background(), pipeline, v1.UpdateOptions{})
	return err
}

func (this *TektonService) DeletePipeline(ns, name string) error {
	return this.Client.TektonV1beta1().Pipelines(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}

func (this *TektonService) ListPipelineRunByNs(ns string) []*v1beta1.PipelineRun {
	return this.PipelineRunMap.ListAll(ns)
}

func (this *TektonService) ShowPipelineRun(ns, name string) *v1beta1.PipelineRun {
	return this.PipelineRunMap.Find(ns, name)
}

func (this *TektonService) CreatePipelineRun(pipelineRun *v1beta1.PipelineRun) error {
	_, err := this.Client.TektonV1beta1().PipelineRuns(pipelineRun.Namespace).Create(context.Background(), pipelineRun, v1.CreateOptions{})
	return err
}

func (this *TektonService) UpdatePipelineRun(pipelineRun *v1beta1.PipelineRun) error {
	_, err := this.Client.TektonV1beta1().PipelineRuns(pipelineRun.Namespace).Update(context.Background(), pipelineRun, v1.UpdateOptions{})
	return err
}

func (this *TektonService) DeletePipelineRun(ns, name string) error {
	return this.Client.TektonV1beta1().PipelineRuns(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}
