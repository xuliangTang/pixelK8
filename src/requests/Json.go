package requests

// CreateIngress 创建ingress对象
type CreateIngress struct {
	Name        string          `json:"name" binding:"required"`
	Namespace   string          `json:"namespace" binding:"required"`
	Rules       []*IngressRules `json:"rules" binding:"required"`
	Annotations string          `json:"annotations"`
}

// IngressRules ingress规则
type IngressRules struct {
	Host  string         `json:"host" binding:"required"`
	Paths []*IngressPath `json:"paths" binding:"required"`
}

// IngressPath ingress配置
type IngressPath struct {
	Path    string `json:"path" binding:"required"`
	SvcName string `json:"svc_name" binding:"required"`
	Port    string `json:"port" binding:"required"`
}