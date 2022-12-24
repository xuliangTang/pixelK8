package dto

type DeploymentList struct {
	Name        string   `json:"name"`
	NameSpace   string   `json:"name_space"`
	Replicas    [3]int32 `json:"replicas"`
	Images      string   `json:"images"`
	IsCompleted bool     `json:"is_completed"`
	Message     string   `json:"message"`
}
