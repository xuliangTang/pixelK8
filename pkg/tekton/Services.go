package tekton

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TektonService struct {
	TaskMap *TaskMapStruct            `inject:"-"`
	Client  *tektonVersiond.Clientset `inject:"-"`
}

func NewTektonService() *TektonService {
	return &TektonService{}
}

func (this *TektonService) LoadTask(ns string) []*v1beta1.Task {
	return this.TaskMap.ListAll(ns)
}

func (this *TektonService) Create(task *v1beta1.Task) error {
	_, err := this.Client.TektonV1beta1().Tasks(task.Namespace).Create(context.Background(), task, v1.CreateOptions{})
	return err
}

func (this *TektonService) Delete(ns, name string) error {
	return this.Client.TektonV1beta1().Tasks(ns).Delete(context.Background(), name, v1.DeleteOptions{})
}
