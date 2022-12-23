package configurations

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/core/informerHandlers"
)

// K8sInformerStart @configuration
type K8sInformerStart struct {
	K8sClient         *kubernetes.Clientset               `inject:"-"`
	DeploymentHandler *informerHandlers.DeploymentHandler `inject:"-"`
}

func NewK8sInformerStart() *K8sInformerStart {
	return &K8sInformerStart{}
}

func (this *K8sInformerStart) InitInformer() informers.SharedInformerFactory {
	informerFactory := informers.NewSharedInformerFactory(this.K8sClient, 0)

	depInformer := informerFactory.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DeploymentHandler)

	informerFactory.Start(wait.NeverStop)

	return informerFactory
}
