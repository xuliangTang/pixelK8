package dto

type NodeList struct {
	Name          string             `json:"name"`
	Ip            string             `json:"ip"`
	Labels        []string           `json:"labels"`
	Taints        []string           `json:"taints"`
	Capacity      *NodeCapacity      `json:"capacity"`
	CapacityUsage *NodeCapacityUsage `json:"capacity_usage"`
	CreatedAt     string             `json:"created_at"`
}

type NodeCapacity struct {
	Cpu    int64 `json:"cpu"`
	Memory int64 `json:"memory"`
	Pods   int64 `json:"pods"`
}

type NodeCapacityUsage struct {
	Cpu    int64 `json:"cpu"`
	Memory int64 `json:"memory"`
	Pods   int64 `json:"pods"`
}
