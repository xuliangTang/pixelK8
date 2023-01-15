package configurations

import (
	"pixelk8/pkg/rbac"
	"pixelk8/src/core/maps"
)

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

func (*K8sMapCfg) InitServiceMap() *maps.ServiceMap {
	return &maps.ServiceMap{}
}

func (*K8sMapCfg) InitSecretMap() *maps.SecretMap {
	return &maps.SecretMap{}
}

func (*K8sMapCfg) InitConfigmapMap() *maps.ConfigmapMap {
	return &maps.ConfigmapMap{}
}

func (*K8sMapCfg) InitNodeMap() *maps.NodeMap {
	return &maps.NodeMap{}
}

func (*K8sMapCfg) InitRoleMap() *rbac.RoleMap {
	return &rbac.RoleMap{}
}

func (*K8sMapCfg) InitRoleBindingMap() *rbac.RoleBindingMap {
	return &rbac.RoleBindingMap{}
}
