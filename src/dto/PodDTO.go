package dto

type PodList struct {
	Name      string `json:"name"`
	NameSpace string `json:"name_space"`
	NodeName  string `json:"node_name"`
	Images    string `json:"images"`
	// 0 podIp 1 hostIp
	Ip [2]string `json:"ip"`
	// 阶段
	Phase string `json:"phase"`
	// pod 是否就绪
	IsReady   bool   `json:"is_ready"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}
