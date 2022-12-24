package maps

import (
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	"sync"
)

type DeploymentMap struct {
	data sync.Map // key:namespace(string), value:[]*appsV1.Deployment
}

// Add 添加
func (this *DeploymentMap) Add(deployment *appsV1.Deployment) {
	if depList, ok := this.data.Load(deployment.Namespace); ok {
		depList = append(depList.([]*appsV1.Deployment), deployment)
		this.data.Store(deployment.Namespace, depList)
	} else {
		this.data.Store(deployment.Namespace, []*appsV1.Deployment{deployment})
	}
}

// ListByNs 获取列表
func (this *DeploymentMap) ListByNs(namespace string) ([]*appsV1.Deployment, error) {
	if depList, ok := this.data.Load(namespace); ok {
		return depList.([]*appsV1.Deployment), nil
	}

	return []*appsV1.Deployment{}, nil
}

// Find 根据名称查找
func (this *DeploymentMap) Find(ns string, depName string) (*appsV1.Deployment, error) {
	if depList, ok := this.data.Load(ns); ok {
		for _, dep := range depList.([]*appsV1.Deployment) {
			if dep.Name == depName {
				return dep, nil
			}
		}
	}

	return &appsV1.Deployment{}, nil
}

// Update 更新
func (this *DeploymentMap) Update(deployment *appsV1.Deployment) error {
	if depList, ok := this.data.Load(deployment.Namespace); ok {
		depList := depList.([]*appsV1.Deployment)
		for i, dep := range depList {
			if dep.Name == deployment.Name {
				depList[i] = deployment
				break
			}
		}
		return nil
	}

	return fmt.Errorf("deployment [%s] not found", deployment.Name)
}

// Delete 删除
func (this *DeploymentMap) Delete(deployment *appsV1.Deployment) {
	if depList, ok := this.data.Load(deployment.Namespace); ok {
		depList := depList.([]*appsV1.Deployment)
		for i, dep := range depList {
			if dep.Name == deployment.Name {
				newDepList := append(depList[:i], depList[i+1:]...)
				this.data.Store(deployment.Namespace, newDepList)
				break
			}
		}
	}
}
