package requests

type DeploymentUri struct {
	Deployment string `uri:"deployment" binding:"required"`
}

type ShowSecretUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"secret" binding:"required"`
}
