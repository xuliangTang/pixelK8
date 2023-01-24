package properties

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
)

var App AppProperties

type AppProperties struct {
	K8s *K8sOpt
}

type K8sOpt struct {
	Host            string
	Port            int32
	DefaultNs       string
	KubeConfigPath  string
	Nodes           map[string]*NodeOpt `mapstructure:"nodes"`
	CACrtPath       string              // CA证书
	CAKeyPath       string              // CA私钥
	UserAccountPath string              // 用户账号证书存储路径
}

type NodeOpt struct {
	Username string
	Password string
	Host     string
	Port     int
}

func (*AppProperties) InitDefaultConfig(viper *viper.Viper) {
	// bind to K8s_NODES env
	viper.BindEnv("k8s.nodes", "K8S_NODES")
	// bind to CA env
	viper.BindEnv("k8s.caCrtPath", "K8S_CRT_PATH")
	viper.BindEnv("k8s.caKeyPath", "K8S_KEY_PATH")
}

// JsonToNodeMapHookFunc 解码nodes字符串
func (*AppProperties) JsonToNodeMapHookFunc() mapstructure.DecodeHookFuncType {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		// Check if the data type matches the expected one
		if f.Kind() != reflect.String {
			return data, nil
		}

		// Check that the target type is our custom type
		if t != reflect.TypeOf(map[string]*NodeOpt{}) {
			return data, nil
		}

		// Format/decode/parse the data and return the new value
		var m map[string]*NodeOpt
		json.Unmarshal([]byte(data.(string)), &m)
		return m, nil
	}
}
