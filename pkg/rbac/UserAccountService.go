package rbac

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"math/big"
	mathRand "math/rand"
	"os"
	"path"
	"path/filepath"
	"pixelk8/src/properties"
	"regexp"
	"strings"
	"time"
)

// UserAccountService @Service
type UserAccountService struct{}

func NewUserAccountService() *UserAccountService {
	return &UserAccountService{}
}

// List 获取userAccount列表
func (*UserAccountService) List() (ret []UserAccountModel, err error) {
	keyReg := regexp.MustCompile(".*_key.pem")
	ret = make([]UserAccountModel, 0)
	const suffix = ".pem"

	err = filepath.Walk(properties.App.K8s.UserAccountPath, func(p string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if path.Ext(f.Name()) == suffix {
			if !keyReg.MatchString(f.Name()) {
				ret = append(ret, UserAccountModel{
					Name:      strings.Replace(f.Name(), suffix, "", -1),
					CreatedAt: f.ModTime().Format(athena.DateTimeFormat),
				})
			}
		}

		return nil
	})

	return ret, err
}

// Paging 分页userAccount切片
func (*UserAccountService) Paging(page *athena.Page, userAccountList []UserAccountModel) athena.Collection {
	count := len(userAccountList)
	iUserAccountList := make([]any, count)

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iUserAccountList)
	collection := athena.NewCollection(userAccountList[start:end], page)
	return *collection
}

// Create 创建userAccount证书和私钥文件
func (this *UserAccountService) Create(cn, o string) error {
	caCert, caPriKey, err := this.ParseK8sCA()
	if err != nil {
		return err
	}

	// 构建证书模板
	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(mathRand.Int63()), // 证书序列号
		Subject: pkix.Name{
			Country:      []string{"CN"},
			Organization: []string{o},
			//OrganizationalUnit: []string{},
			Province:   []string{cn},
			CommonName: cn, // CN
			Locality:   []string{cn},
		},
		NotBefore:             time.Now(),                                                                 // 证书有效期开始时间
		NotAfter:              time.Now().AddDate(1, 0, 0),                                                // 证书有效期
		BasicConstraintsValid: true,                                                                       // 基本的有效性约束
		IsCA:                  false,                                                                      // 是否是根证书
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, // 证书用途(客户端认证，数据加密)
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		EmailAddresses:        []string{},
	}

	// 生成公私钥秘钥对
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// 创建证书对象
	clientCert, err := x509.CreateCertificate(rand.Reader, certTemplate, caCert, &priKey.PublicKey, caPriKey)
	if err != nil {
		return err
	}

	// 编码证书文件和私钥文件
	clientCertPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientCert,
	}

	clientCertFile, err := os.OpenFile(fmt.Sprintf("%s/%s.pem", properties.App.K8s.UserAccountPath, cn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	err = pem.Encode(clientCertFile, clientCertPem)
	if err != nil {
		return err
	}

	buf := x509.MarshalPKCS1PrivateKey(priKey)
	keyPem := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: buf,
	}
	clientKeyFile, _ := os.OpenFile(fmt.Sprintf("%s/%s_key.pem", properties.App.K8s.UserAccountPath, cn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	err = pem.Encode(clientKeyFile, keyPem)
	return err
}

// Delete 删除userAccount证书和私钥文件
func (*UserAccountService) Delete(cn string) error {
	err := os.Remove(fmt.Sprintf("%s/%s.pem", properties.App.K8s.UserAccountPath, cn))
	if err != nil {
		return err
	}

	err = os.Remove(fmt.Sprintf("%s/%s_key.pem", properties.App.K8s.UserAccountPath, cn))
	return err
}

// ParseK8sCA 解析k8s CA证书
func (*UserAccountService) ParseK8sCA() (*x509.Certificate, *rsa.PrivateKey, error) {
	// 解析证书
	caFile, err := os.ReadFile(properties.App.K8s.CACrtPath)
	if err != nil {
		return nil, nil, err
	}
	caBlock, _ := pem.Decode(caFile)
	caCert, err := x509.ParseCertificate(caBlock.Bytes) // CA证书对象
	if err != nil {
		return nil, nil, err
	}

	// 解析私钥
	keyFile, err := os.ReadFile(properties.App.K8s.CAKeyPath)
	if err != nil {
		return nil, nil, err
	}
	keyBlock, _ := pem.Decode(keyFile)
	caPriKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes) // 私钥对象
	if err != nil {
		return nil, nil, err
	}

	return caCert, caPriKey, nil
}

// Kubeconfig 生成userAccount的kubeconfig
func (*UserAccountService) Kubeconfig(cn string) ([]byte, error) {
	cfg := api.NewConfig()
	const clusterName = "kubernetes"

	k8sCertData, err := CertData(properties.App.K8s.CACrtPath)
	if err != nil {
		return nil, err
	}
	cfg.Clusters[clusterName] = &api.Cluster{
		Server:                   properties.App.K8s.ApiServer,
		CertificateAuthorityData: k8sCertData,
	}

	contextName := fmt.Sprintf("%s@kubernetes", cn)
	cfg.Contexts[contextName] = &api.Context{
		AuthInfo: cn,
		Cluster:  clusterName,
	}
	cfg.CurrentContext = contextName

	userCertFile := fmt.Sprintf("%s/%s.pem", properties.App.K8s.UserAccountPath, cn)
	userCertKeyFile := fmt.Sprintf("%s/%s_key.pem", properties.App.K8s.UserAccountPath, cn)
	clientKeyData, err := CertData(userCertKeyFile)
	if err != nil {
		return nil, err
	}
	clientCertData, err := CertData(userCertFile)
	if err != nil {
		return nil, err
	}
	cfg.AuthInfos[cn] = &api.AuthInfo{
		ClientKeyData:         clientKeyData,
		ClientCertificateData: clientCertData,
	}

	fileContent, err := clientcmd.Write(*cfg)
	return fileContent, err
}
