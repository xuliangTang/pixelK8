package dto

import coreV1 "k8s.io/api/core/v1"

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

type NodeShow struct {
	Name      string                `json:"name"`
	Ip        string                `json:"ip"`
	Labels    map[string]string     `json:"labels"`
	Taints    []coreV1.Taint        `json:"taints"`
	Info      coreV1.NodeSystemInfo `json:"info"`
	CreatedAt string                `json:"created_at"`
}
