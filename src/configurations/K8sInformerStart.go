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
	ReplicaSetHandler *informerHandlers.ReplicaSetHandler `inject:"-"`
	PodHandler        *informerHandlers.PodHandler        `inject:"-"`
	EventHandler      *informerHandlers.EventHandler      `inject:"-"`
}

func NewK8sInformerStart() *K8sInformerStart {
	return &K8sInformerStart{}
}

func (this *K8sInformerStart) InitInformer() informers.SharedInformerFactory {
	informerFactory := informers.NewSharedInformerFactory(this.K8sClient, 0)

	depInformer := informerFactory.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DeploymentHandler)

	rsInformer := informerFactory.Apps().V1().ReplicaSets()
	rsInformer.Informer().AddEventHandler(this.ReplicaSetHandler)

	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(this.PodHandler)

	eventInformer := informerFactory.Core().V1().Events()
	eventInformer.Informer().AddEventHandler(this.EventHandler)

	informerFactory.Start(wait.NeverStop)

	return informerFactory
}
