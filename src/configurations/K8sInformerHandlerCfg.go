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
