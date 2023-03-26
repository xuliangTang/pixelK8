package configurations

import (
	tektonVersiond "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

func (*K8sCfg) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", properties.App.K8s.KubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

// InitClient 初始化 k8s client
func (this *K8sCfg) InitClient() *kubernetes.Clientset {
	/*config := &rest.Config{
		Host:        fmt.Sprintf("%s:%d", properties.App.K8s.Host, properties.App.K8s.Port),
		BearerToken: "",
	}*/

	client, err := kubernetes.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// InitMetricsClient 初始化 metrics client
func (this *K8sCfg) InitMetricsClient() *versioned.Clientset {
	client, err := versioned.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// InitTektonClient 初始化tekton客户端
func (this *K8sCfg) InitTektonClient() *tektonVersiond.Clientset {
	client, err := tektonVersiond.NewForConfig(this.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
