package dto

type NodeList struct {
	Name      string   `json:"name"`
	Ip        string   `json:"ip"`
	Labels    []string `json:"labels"`
	CreatedAt string   `json:"created_at"`
}
