package dto

type SecretList struct {
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Type      [2]string `json:"type"`
	CreatedAt string    `json:"created_at"`
}
