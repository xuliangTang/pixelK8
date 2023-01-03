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

// CreateSecret 创建secret对象
type CreateSecret struct {
	Name      string            `json:"name" binding:"required"`
	Namespace string            `json:"namespace" binding:"required"`
	Type      string            `json:"type" binding:"required"`
	Data      map[string]string `json:"data" binding:"required"`
	IsEdit    *bool             `json:"is_edit"`
}

// CreateConfigmap 创建configmap对象
type CreateConfigmap struct {
	Name      string            `json:"name" binding:"required"`
	Namespace string            `json:"namespace" binding:"required"`
	Data      map[string]string `json:"data" binding:"required"`
	IsEdit    *bool             `json:"is_edit"`
}
