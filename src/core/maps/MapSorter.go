package maps

import (
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	networkingV1 "k8s.io/api/networking/v1"
)

// 排序namespaces 按名称
type coreV1Namespace []*coreV1.Namespace

func (this coreV1Namespace) Len() int {
	return len(this)
}

func (this coreV1Namespace) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this coreV1Namespace) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 排序pods 按创建时间
type coreV1Pods []*coreV1.Pod

func (this coreV1Pods) Len() int {
	return len(this)
}

func (this coreV1Pods) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this coreV1Pods) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 排序deployments 按创建时间
type appsV1Deployments []*appsV1.Deployment

func (this appsV1Deployments) Len() int {
	return len(this)
}

func (this appsV1Deployments) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this appsV1Deployments) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 排序ingress 按创建时间
type networkingV1Ingress []*networkingV1.Ingress

func (this networkingV1Ingress) Len() int {
	return len(this)
}

func (this networkingV1Ingress) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this networkingV1Ingress) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 排序service 按创建时间
type coreV1Service []*coreV1.Service

func (this coreV1Service) Len() int {
	return len(this)
}

func (this coreV1Service) Less(i, j int) bool {
	return this[i].CreationTimestamp.Before(&this[j].CreationTimestamp)
}

func (this coreV1Service) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
