package properties

import "github.com/spf13/viper"

var App AppProperties

type AppProperties struct {
	K8s *K8sOpt
}

type K8sOpt struct {
	Host string
	Port int32
}

func (*AppProperties) InitDefaultConfig(viper *viper.Viper) {}
