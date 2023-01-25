package requests

type DeploymentUri struct {
	Deployment string `uri:"deployment" binding:"required"`
}

type ShowDeploymentUri struct {
	Namespace  string `uri:"ns" binding:"required"`
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

type PodContainersLogsUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"pod" binding:"required"`
}

type DeleteDeploymentUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"deployment" binding:"required"`
}

type DeletePodUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"pod" binding:"required"`
}

type DeleteIngressUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"ingress" binding:"required"`
}

type DeleteSecretUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"secret" binding:"required"`
}

type DeleteConfigmapUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"configmap" binding:"required"`
}

type DeleteServiceUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"service" binding:"required"`
}

type PodContainerTerminalUri struct {
	Namespace string `uri:"ns" binding:"required"`
	Name      string `uri:"pod" binding:"required"`
}

type NodeTerminalUri struct {
	Name string `uri:"node" binding:"required"`
}

type ShowNodeUri struct {
	Name string `uri:"node" binding:"required"`
}
