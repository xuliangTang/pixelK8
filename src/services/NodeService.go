package services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"golang.org/x/crypto/ssh"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"net"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/requests"
	"regexp"
)

const hostPattern = "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"

// NodeService @Service
type NodeService struct {
	NodeMap       *maps.NodeMap         `inject:"-"`
	PodMap        *maps.PodMap          `inject:"-"`
	MetricsClient *versioned.Clientset  `inject:"-"`
	K8sClient     *kubernetes.Clientset `inject:"-"`
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

// List node列表
func (this *NodeService) List() (ret []*dto.NodeList) {
	nodeList := this.NodeMap.List()

	ret = make([]*dto.NodeList, len(nodeList))
	for i, node := range nodeList {
		ret[i] = &dto.NodeList{
			Name:   node.Name,
			Ip:     node.Status.Addresses[0].Address,
			Labels: this.filterLabels(node.Labels),
			Taints: this.filterTaints(node.Spec.Taints),
			Capacity: &dto.NodeCapacity{
				Cpu:    node.Status.Capacity.Cpu().Value(),
				Memory: node.Status.Capacity.Memory().Value(),
				Pods:   node.Status.Capacity.Pods().Value(),
			},
			CapacityUsage: this.getMetricsCapacityUsageByNodeName(node.Name),
			CreatedAt:     node.CreationTimestamp.Format(athena.DateTimeFormat),
		}
	}

	return
}

// Paging 分页node切片
func (this *NodeService) Paging(page *athena.Page, nodeList []*dto.NodeList) athena.Collection {
	var count int
	iNodeList := make([]any, len(nodeList))
	for i, node := range iNodeList {
		count++
		iNodeList[i] = node
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iNodeList)
	collection := athena.NewCollection(nodeList[start:end], page)
	return *collection
}

// Show 获取node详情
func (this *NodeService) Show(uri *requests.ShowNodeUri) *dto.NodeShow {
	node := this.NodeMap.Find(uri.Name)
	return &dto.NodeShow{
		Name:      node.Name,
		Ip:        node.Status.Addresses[0].Address,
		Labels:    node.Labels,
		Taints:    node.Spec.Taints,
		Info:      node.Status.NodeInfo,
		CreatedAt: node.CreationTimestamp.Format(athena.DateTimeFormat),
	}
}

// Update 更新node
func (this *NodeService) Update(uri *requests.ShowNodeUri, req *requests.UpdateNode) error {
	node := this.NodeMap.Find(uri.Name)
	node.Labels = req.Labels
	node.Spec.Taints = req.Taints

	_, err := this.K8sClient.CoreV1().Nodes().Update(context.Background(), node, metaV1.UpdateOptions{})
	return err
}

// SSHConnect 获取SSH连接
func (*NodeService) SSHConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallback,
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

// 过滤域名形式的labels
func (*NodeService) filterLabels(labels map[string]string) (ret []string) {
	ret = make([]string, 0)
	for k, v := range labels {
		if !regexp.MustCompile(hostPattern).MatchString(k) {
			ret = append(ret, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return
}

// 过滤域名形式的taints
func (*NodeService) filterTaints(taints []coreV1.Taint) (ret []string) {
	ret = make([]string, 0)
	for _, taint := range taints {
		if !regexp.MustCompile(hostPattern).MatchString(taint.Key) {
			ret = append(ret, fmt.Sprintf("%s=%s:%s", taint.Key, taint.Value, taint.Effect))
		}
	}
	return
}

// 通过metrics获取节点已使用资源
func (this *NodeService) getMetricsCapacityUsageByNodeName(nodeName string) *dto.NodeCapacityUsage {
	nodeMetrics, err := this.MetricsClient.MetricsV1beta1().NodeMetricses().Get(context.Background(), nodeName, metaV1.GetOptions{})
	athena.Error(err)

	pods := this.PodMap.CountByNodeName(nodeName)
	return &dto.NodeCapacityUsage{
		Cpu:    nodeMetrics.Usage.Cpu().MilliValue(),
		Memory: nodeMetrics.Usage.Memory().MilliValue(),
		Pods:   pods,
	}
}
