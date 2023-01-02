package requests

type PodContainerLogsQuery struct {
	ContainerName string `form:"container_name" binding:"required"`
}
