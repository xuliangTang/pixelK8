package configurations

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"pixelk8/src/properties"
)

// K8sCfg @configuration
type K8sCfg struct{}

func NewK8sCfg() *K8sCfg {
	return &K8sCfg{}
}

// InitClient 初始化 k8s client
func (*K8sCfg) InitClient() *kubernetes.Clientset {
	/*config := &rest.Config{
		Host:        fmt.Sprintf("%s:%d", properties.App.K8s.Host, properties.App.K8s.Port),
		BearerToken: "",
	}*/

	config, err := clientcmd.BuildConfigFromFlags("", properties.App.K8s.KubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// InitMetricsClient 初始化 metrics client
func (*K8sCfg) InitMetricsClient() *versioned.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", properties.App.K8s.KubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
