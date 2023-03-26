package configurations

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"pixelk8/pkg/rbac"
	"pixelk8/pkg/tekton"
	"pixelk8/src/core/informerHandlers"
	"pixelk8/src/properties"
)

// K8sInformerStart @configuration
type K8sInformerStart struct {
	K8sClient                 *kubernetes.Clientset               `inject:"-"`
	DeploymentHandler         *informerHandlers.DeploymentHandler `inject:"-"`
	ReplicaSetHandler         *informerHandlers.ReplicaSetHandler `inject:"-"`
	PodHandler                *informerHandlers.PodHandler        `inject:"-"`
	EventHandler              *informerHandlers.EventHandler      `inject:"-"`
	NamespaceHandler          *informerHandlers.NamespaceHandler  `inject:"-"`
	IngressHandler            *informerHandlers.IngressHandler    `inject:"-"`
	ServiceHandler            *informerHandlers.ServiceHandler    `inject:"-"`
	SecretHandler             *informerHandlers.SecretHandler     `inject:"-"`
	ConfigmapHandler          *informerHandlers.ConfigmapHandler  `inject:"-"`
	NodeHandler               *informerHandlers.NodeHandler       `inject:"-"`
	RoleHandler               *rbac.RoleHandler                   `inject:"-"`
	RoleBindingHandler        *rbac.RoleBindingHandler            `inject:"-"`
	ServiceAccountHandler     *rbac.ServiceAccountHandler         `inject:"-"`
	ClusterRoleHandler        *rbac.ClusterRoleHandler            `inject:"-"`
	ClusterRoleBindingHandler *rbac.ClusterRoleBindingHandler     `inject:"-"`
	TektonTaskHandler         *tekton.TaskHandler                 `inject:"-"`
	TektonPipelineHandler     *tekton.PipelineHandler             `inject:"-"`
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

	nsInformer := informerFactory.Core().V1().Namespaces()
	nsInformer.Informer().AddEventHandler(this.NamespaceHandler)

	ingInformer := informerFactory.Networking().V1().Ingresses()
	ingInformer.Informer().AddEventHandler(this.IngressHandler)

	svcInformer := informerFactory.Core().V1().Services()
	svcInformer.Informer().AddEventHandler(this.ServiceHandler)

	secretInformer := informerFactory.Core().V1().Secrets()
	secretInformer.Informer().AddEventHandler(this.SecretHandler)

	configmapInformer := informerFactory.Core().V1().ConfigMaps()
	configmapInformer.Informer().AddEventHandler(this.ConfigmapHandler)

	nodeInformer := informerFactory.Core().V1().Nodes()
	nodeInformer.Informer().AddEventHandler(this.NodeHandler)

	roleInformer := informerFactory.Rbac().V1().Roles()
	roleInformer.Informer().AddEventHandler(this.RoleHandler)

	roleBindingInformer := informerFactory.Rbac().V1().RoleBindings()
	roleBindingInformer.Informer().AddEventHandler(this.RoleBindingHandler)

	serviceAccountInformer := informerFactory.Core().V1().ServiceAccounts()
	serviceAccountInformer.Informer().AddEventHandler(this.ServiceAccountHandler)

	clusterRoleInformer := informerFactory.Rbac().V1().ClusterRoles()
	clusterRoleInformer.Informer().AddEventHandler(this.ClusterRoleHandler)

	clusterRoleBindingInformer := informerFactory.Rbac().V1().ClusterRoleBindings()
	clusterRoleBindingInformer.Informer().AddEventHandler(this.ClusterRoleBindingHandler)

	informerFactory.Start(wait.NeverStop)

	return informerFactory
}

func (*K8sInformerStart) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", properties.App.K8s.KubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

// 动态informers处理
var taskResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "tasks", Version: "v1beta1"}
var pipelineResource = schema.GroupVersionResource{Group: "tekton.dev", Resource: "pipelines", Version: "v1beta1"}

func (this *K8sInformerStart) InitGenericInformer() dynamicinformer.DynamicSharedInformerFactory {
	client, err := dynamic.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	di := dynamicinformer.NewDynamicSharedInformerFactory(client, 0)

	taskInformer := di.ForResource(taskResource)
	taskInformer.Informer().AddEventHandler(this.TektonTaskHandler)
	di.ForResource(pipelineResource).Informer().AddEventHandler(this.TektonPipelineHandler)

	di.Start(wait.NeverStop)
	return di
}
