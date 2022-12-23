package configurations

import "pixelk8/src/core/maps"

// K8sMapCfg @configuration
type K8sMapCfg struct{}

func NewK8sMapCfg() *K8sMapCfg {
	return &K8sMapCfg{}
}

func (*K8sMapCfg) InitDeploymentMap() *maps.DeploymentMap {
	return &maps.DeploymentMap{}
}
