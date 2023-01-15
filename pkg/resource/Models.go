package resource

type GroupResource struct {
	Group     string     `json:"group"`
	Version   string     `json:"version"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Name  string   `json:"name"`
	Verbs []string `json:"verbs"`
}
