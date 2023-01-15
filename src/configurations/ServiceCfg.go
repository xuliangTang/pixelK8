package configurations

import (
	"pixelk8/pkg/rbac"
	"pixelk8/src/services"
)

// ServiceCfg @configuration
type ServiceCfg struct{}

func NewServiceCfg() *ServiceCfg {
	return &ServiceCfg{}
}

func (*ServiceCfg) InitCommonService() *services.CommonService {
	return services.NewCommonService()
}

func (*ServiceCfg) InitDeploymentService() *services.DeploymentService {
	return services.NewDeploymentService()
}

func (*ServiceCfg) InitPodService() *services.PodService {
	return services.NewPodService()
}

func (*ServiceCfg) InitNamespaceService() *services.NamespaceService {
	return services.NewNamespaceService()
}

func (*ServiceCfg) InitIngressService() *services.IngressService {
	return services.NewIngressService()
}

func (*ServiceCfg) InitServiceService() *services.ServiceService {
	return services.NewServiceService()
}

func (*ServiceCfg) InitSecretService() *services.SecretService {
	return services.NewSecretService()
}

func (*ServiceCfg) InitConfigmapService() *services.ConfigmapService {
	return services.NewConfigmapService()
}

func (*ServiceCfg) InitNodeService() *services.NodeService {
	return services.NewNodeService()
}

func (*ServiceCfg) InitRoleService() *rbac.RoleService {
	return rbac.NewRoleService()
}

func (*ServiceCfg) InitRoleBindingService() *rbac.RoleBindingService {
	return rbac.NewRoleBindingService()
}
