package properties

import "github.com/spf13/viper"

var App AppProperties

type AppProperties struct {
	K8s *K8sOpt
}

type K8sOpt struct {
	Host           string
	Port           int32
	DefaultNs      string
	KubeConfigPath string
	Nodes          map[string]*NodeOpt `mapstructure:"nodes"`
}

type NodeOpt struct {
	Username string
	Password string
	Host     string
	Port     int
}

func (*AppProperties) InitDefaultConfig(viper *viper.Viper) {}
