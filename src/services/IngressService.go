package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	networkingV1 "k8s.io/api/networking/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"pixelk8/src/constants"
	"pixelk8/src/core/maps"
	"pixelk8/src/dto"
	"pixelk8/src/requests"
	"sigs.k8s.io/yaml"
	"strconv"
)

// IngressService @Service
type IngressService struct {
	IngressMap *maps.IngressMap      `inject:"-"`
	K8sClient  *kubernetes.Clientset `inject:"-"`
	CommonSvc  *CommonService        `inject:"-"`
}

func NewIngressService() *IngressService {
	return &IngressService{}
}

// ListByNs 根据ns获取ingress列表
func (this *IngressService) ListByNs(ns string) (ret []*dto.IngressList) {
	ingList := this.IngressMap.ListByNs(ns)

	ret = make([]*dto.IngressList, len(ingList))
	for i, ing := range ingList {
		ret[i] = &dto.IngressList{
			Name:      ing.Name,
			Namespace: ing.Namespace,
			Hosts:     this.joinRuleHosts(ing.Spec.Rules),
			CreatedAt: ing.CreationTimestamp.Format(athena.DateTimeFormat),
			Opt: &dto.IngressOpt{
				CorsEnable:    this.checkOpt(ing, constants.CorsEnable),
				RewriteEnable: this.checkOpt(ing, constants.RewriteEnable),
			},
		}
	}

	return
}

// Paging 分页ingress切片
func (this *IngressService) Paging(page *athena.Page, ingList []*dto.IngressList) athena.Collection {
	var count int
	iIngList := make([]any, len(ingList))
	for i, ing := range ingList {
		count++
		iIngList[i] = ing
	}

	page.Extend = gin.H{"count": count}
	// 分页
	start, end := page.SlicePage(iIngList)
	collection := athena.NewCollection(ingList[start:end], page)
	return *collection
}

// Create 创建ingress
func (this *IngressService) Create(req *requests.CreateIngress) error {
	className := "nginx"
	var ingressRules []networkingV1.IngressRule
	prefix := networkingV1.PathTypePrefix

	for _, r := range req.Rules {
		httpRuleValue := &networkingV1.HTTPIngressRuleValue{}
		rulePaths := make([]networkingV1.HTTPIngressPath, 0)

		for _, pathCfg := range r.Paths {
			port, err := strconv.Atoi(pathCfg.Port)
			if err != nil {
				return err
			}
			rulePaths = append(rulePaths, networkingV1.HTTPIngressPath{
				Path:     pathCfg.Path,
				PathType: &prefix,
				Backend: networkingV1.IngressBackend{
					Service: &networkingV1.IngressServiceBackend{
						Name: pathCfg.SvcName,
						Port: networkingV1.ServiceBackendPort{
							Number: int32(port),
						},
					},
				},
			})
		}

		httpRuleValue.Paths = rulePaths
		rule := networkingV1.IngressRule{
			Host: r.Host,
			IngressRuleValue: networkingV1.IngressRuleValue{
				HTTP: httpRuleValue,
			},
		}
		ingressRules = append(ingressRules, rule)
	}

	ingress := &networkingV1.Ingress{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:        req.Name,
			Namespace:   req.Namespace,
			Annotations: this.CommonSvc.ParseAnnotations(req.Annotations),
		},
		Spec: networkingV1.IngressSpec{
			IngressClassName: &className,
			Rules:            ingressRules,
		},
	}

	_, err := this.K8sClient.NetworkingV1().Ingresses(req.Namespace).
		Create(context.Background(), ingress, metaV1.CreateOptions{})

	return err
}

// Delete 删除ingress
func (this *IngressService) Delete(uri *requests.DeleteIngressUri) error {
	return this.K8sClient.NetworkingV1().Ingresses(uri.Namespace).
		Delete(context.Background(), uri.Name, metaV1.DeleteOptions{})
}

// GetForYaml 查询ingress转为yaml
func (this *IngressService) GetForYaml(uri *requests.NamespaceNameUri) ([]byte, error) {
	ing, err := this.K8sClient.NetworkingV1().Ingresses(uri.Namespace).Get(context.Background(), uri.Name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	b, err := yaml.Marshal(ing)
	return b, err
}

// 拼接ingress host
func (*IngressService) joinRuleHosts(rules []networkingV1.IngressRule) (ret []string) {
	ret = make([]string, 0)
	for _, rule := range rules {
		for _, path := range rule.HTTP.Paths {
			ret = append(ret, rule.Host+path.Path)
		}
	}

	return
}

// 判断ingress是否开启opt
func (*IngressService) checkOpt(ingress *networkingV1.Ingress, opt string) bool {
	if _, ok := ingress.Annotations[opt]; ok {
		return true
	}

	return false
}
