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

func (*K8sInformerHandlerCfg) EventHandler() *informerHandlers.EventHandler {
	return &informerHandlers.EventHandler{}
}
