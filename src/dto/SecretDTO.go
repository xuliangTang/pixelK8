package dto

type SecretList struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	CreatedAt string `json:"created_at"`
}
