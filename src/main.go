package main

import (
	"github.com/xuliangTang/athena/athena"
	"pixelk8/src/configurations"
	"pixelk8/src/controllers"
	"pixelk8/src/properties"
)

func main() {
	athena.Ignite().
		MappingConfig(&properties.App).
		Configuration(
			configurations.NewK8sMapCfg(),
			configurations.NewK8sCfg(),
			configurations.NewK8sInformerHandlerCfg(),
			configurations.NewK8sInformerStart(),
			configurations.NewServiceCfg()).
		Mount("v1", nil, controllers.NewDeploymentCtl()).
		Launch()
}
