package dto

type ServiceList struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Type      string   `json:"type"`
	Selector  []string `json:"selector"`
	Target    []string `json:"target"`
	CreatedAt string   `json:"created_at"`
}
