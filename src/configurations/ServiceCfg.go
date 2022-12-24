package configurations

import "pixelk8/src/services"

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
