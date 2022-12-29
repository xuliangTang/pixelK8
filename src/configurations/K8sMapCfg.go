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

func (*K8sMapCfg) InitReplicaSetMap() *maps.ReplicaSetMap {
	return &maps.ReplicaSetMap{}
}

func (*K8sMapCfg) InitPodMap() *maps.PodMap {
	return &maps.PodMap{}
}

func (*K8sMapCfg) InitEventMap() *maps.EventMap {
	return &maps.EventMap{}
}

func (*K8sMapCfg) InitNamespaceMap() *maps.NamespaceMap {
	return &maps.NamespaceMap{}
}

func (*K8sMapCfg) InitIngressMap() *maps.IngressMap {
	return &maps.IngressMap{}
}
