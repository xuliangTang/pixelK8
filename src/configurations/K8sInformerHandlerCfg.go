package configurations

import (
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
