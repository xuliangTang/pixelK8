package requests

type DeploymentUri struct {
	Deployment string `uri:"deployment" binding:"required"`
}
