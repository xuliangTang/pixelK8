package configurations

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
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

	config, err := clientcmd.BuildConfigFromFlags("", "kubeconfig")
	if err != nil {
		log.Fatal(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
