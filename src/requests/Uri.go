package requests

type DeploymentUri struct {
	Deployment string `uri:"deployment" binding:"required"`
}

type ShowSecretUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"secret" binding:"required"`
}

type ShowConfigmapUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"configmap" binding:"required"`
}

type PodAllContainersUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"pod" binding:"required"`
}
