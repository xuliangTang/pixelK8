package dto

type IngressList struct {
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
	Hosts     []string    `json:"hosts"`
	CreatedAt string      `json:"created_at"`
	Opt       *IngressOpt `json:"opt"`
}

type IngressOpt struct {
	CorsEnable    bool `json:"cors_enable"`
	RewriteEnable bool `json:"rewrite_enable"`
}
