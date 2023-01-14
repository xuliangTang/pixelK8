package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"golang.org/x/crypto/ssh"
	"net"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"regexp"
)

// NodeService @Service
type NodeService struct {
	NodeMap *maps.NodeMap `inject:"-"`
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
			Name:      node.Name,
			Ip:        node.Status.Addresses[0].Address,
			Labels:    this.filterLabels(node.Labels),
			CreatedAt: node.CreationTimestamp.Format(athena.DateTimeFormat),
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
	const expr = "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"
	ret = make([]string, 0)
	for k, v := range labels {
		if !regexp.MustCompile(expr).MatchString(k) {
			ret = append(ret, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return
}
