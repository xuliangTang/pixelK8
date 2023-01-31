package requests

type PodContainerLogsQuery struct {
	ContainerName string `form:"container_name" binding:"required"`
}

type PodContainerTerminalQuery struct {
	ContainerName string `form:"container_name" binding:"required"`
}

type CreateDeploymentQuery struct {
	Fastmod bool `form:"fastmod"`
}
