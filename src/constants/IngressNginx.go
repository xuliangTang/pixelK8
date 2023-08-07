package constants

const (
	CorsEnable    = "nginx.ingress.kubernetes.io/enable-cors"
	RewriteEnable = "nginx.ingress.kubernetes.io/rewrite-enable"
	AuthEnable    = "nginx.ingress.kubernetes.io/auth-enable"
	LimitEnable   = "nginx.ingress.kubernetes.io/limit-enable"
	CanaryEnable  = "nginx.ingress.kubernetes.io/canary"
	MirrorEnable  = "nginx.ingress.kubernetes.io/mirror-enable"
)
