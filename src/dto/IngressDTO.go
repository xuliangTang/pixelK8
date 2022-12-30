package dto

type IngressList struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Hosts     []string `json:"hosts"`
	CreatedAt string   `json:"created_at"`
}
