package dto

type DeploymentList struct {
	Name        string   `json:"name"`
	Namespace   string   `json:"namespace"`
	Replicas    [3]int32 `json:"replicas"`
	Images      string   `json:"images"`
	IsCompleted bool     `json:"is_completed"`
	Message     string   `json:"message"`
	CreatedAt   string   `json:"created_at"`
}
