package main

import (
	"github.com/spf13/viper"
	"github.com/xuliangTang/athena/athena"
	"github.com/xuliangTang/athena/athena/plugins"
	"pixelk8/pkg/rbac"
	"pixelk8/pkg/resource"
	"pixelk8/src/configurations"
	"pixelk8/src/controllers"
	"pixelk8/src/properties"
)

func main() {
	server := athena.Ignite().
		MappingConfig(&properties.App, viper.DecodeHook(properties.App.JsonToNodeMapHookFunc())).
		RegisterPlugin(plugins.NewI18n()).
		Configuration(
			configurations.NewK8sMapCfg(),
			configurations.NewK8sCfg(),
			configurations.NewK8sInformerHandlerCfg(),
			configurations.NewK8sInformerStart(),
			configurations.NewServiceCfg(),
			configurations.NewLocalizeCfg()).
		Mount("v1", nil,
			controllers.NewDeploymentCtl(),
			controllers.NewWsCtl(),
			controllers.NewNamespaceCtl(),
			controllers.NewPodCtl(),
			controllers.NewIngressCtl(),
			controllers.NewServiceCtl(),
			controllers.NewSecretCtl(),
			controllers.NewConfigmapCtl(),
			controllers.NewNodeCtl(),
			resource.NewResourceCtl(),
			rbac.NewRBACCtl())

	/*server.GET("/c/*filepath", func(context *gin.Context) {
		http.FileServer(FS(false)).ServeHTTP(context.Writer, context.Request)
	})*/

	server.Launch()
}
