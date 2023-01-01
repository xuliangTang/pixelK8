package dto

type SecretList struct {
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Type      [2]string `json:"type"`
	CreatedAt string    `json:"created_at"`
}

type SecretShow struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Type      [2]string         `json:"type"`
	CreatedAt string            `json:"created_at"`
	Data      map[string][]byte `json:"data"`
	Ext       any               `json:"ext"`
}
