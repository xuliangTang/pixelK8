package configurations

import (
	"pixelk8/pkg/rbac"
	"pixelk8/pkg/tekton"
	"pixelk8/src/core/informerHandlers"
)

// K8sInformerHandlerCfg @configuration
type K8sInformerHandlerCfg struct{}

func NewK8sInformerHandlerCfg() *K8sInformerHandlerCfg {
	return &K8sInformerHandlerCfg{}
}

func (*K8sInformerHandlerCfg) InitDeploymentHandler() *informerHandlers.DeploymentHandler {
	return &informerHandlers.DeploymentHandler{}
}

func (*K8sInformerHandlerCfg) InitReplicaSetHandler() *informerHandlers.ReplicaSetHandler {
	return &informerHandlers.ReplicaSetHandler{}
}

func (*K8sInformerHandlerCfg) InitPodHandler() *informerHandlers.PodHandler {
	return &informerHandlers.PodHandler{}
}

func (*K8sInformerHandlerCfg) InitEventHandler() *informerHandlers.EventHandler {
	return &informerHandlers.EventHandler{}
}

func (*K8sInformerHandlerCfg) InitNamespaceHandler() *informerHandlers.NamespaceHandler {
	return &informerHandlers.NamespaceHandler{}
}

func (*K8sInformerHandlerCfg) InitIngressHandler() *informerHandlers.IngressHandler {
	return &informerHandlers.IngressHandler{}
}

func (*K8sInformerHandlerCfg) InitServiceHandler() *informerHandlers.ServiceHandler {
	return &informerHandlers.ServiceHandler{}
}

func (*K8sInformerHandlerCfg) InitSecretHandler() *informerHandlers.SecretHandler {
	return &informerHandlers.SecretHandler{}
}

func (*K8sInformerHandlerCfg) InitConfigmapHandler() *informerHandlers.ConfigmapHandler {
	return &informerHandlers.ConfigmapHandler{}
}

func (*K8sInformerHandlerCfg) InitNodeHandler() *informerHandlers.NodeHandler {
	return &informerHandlers.NodeHandler{}
}

func (*K8sInformerHandlerCfg) InitRoleHandler() *rbac.RoleHandler {
	return &rbac.RoleHandler{}
}

func (*K8sInformerHandlerCfg) InitRoleBindingHandler() *rbac.RoleBindingHandler {
	return &rbac.RoleBindingHandler{}
}

func (*K8sInformerHandlerCfg) InitServiceAccountHandler() *rbac.ServiceAccountHandler {
	return &rbac.ServiceAccountHandler{}
}

func (*K8sInformerHandlerCfg) InitClusterRoleHandler() *rbac.ClusterRoleHandler {
	return &rbac.ClusterRoleHandler{}
}

func (*K8sInformerHandlerCfg) InitClusterRoleBindingHandler() *rbac.ClusterRoleBindingHandler {
	return &rbac.ClusterRoleBindingHandler{}
}

func (*K8sInformerHandlerCfg) InitTektonTaskHandler() *tekton.TaskHandler {
	return &tekton.TaskHandler{}
}
