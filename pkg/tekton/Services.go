package tekton

import "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

type TektonService struct {
	TaskMap *TaskMapStruct `inject:"-"`
}

func NewTektonService() *TektonService {
	return &TektonService{}
}

func (this *TektonService) LoadTask(ns string) []*v1beta1.Task {
	return this.TaskMap.ListAll(ns)
}
