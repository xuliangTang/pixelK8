package dto

type ConfigmapList struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	CreatedAt string `json:"created_at"`
}

type ConfigmapShow struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	CreatedAt string            `json:"created_at"`
	Data      map[string]string `json:"data"`
}
