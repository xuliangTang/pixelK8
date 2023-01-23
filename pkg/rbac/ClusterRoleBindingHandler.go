package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/tools/cache"
	"pixelk8/src/dto"
	"pixelk8/src/ws"
)

type ClusterRoleBindingHandler struct {
	ClusterRoleBindingMap *ClusterRoleBindingMap     `inject:"-"`
	ClusterRoleBindingSvc *ClusterRoleBindingService `inject:"-"`
}

func (this *ClusterRoleBindingHandler) OnAdd(obj interface{}) {
	if clusterRoleBinding, ok := obj.(*rbacV1.ClusterRoleBinding); ok {
		this.ClusterRoleBindingMap.Add(clusterRoleBinding)

		// 通知ws客户端
		clusterRoleBindingList := this.ClusterRoleBindingSvc.List()
		page := athena.NewPage(1, 10)
		ws.ClientMap.SendAll(dto.WS{
			Type: "clusterRoleBindings",
			Result: gin.H{
				"ns":   "",
				"data": this.ClusterRoleBindingSvc.Paging(page, clusterRoleBindingList),
			},
		})
	}
}

func (this *ClusterRoleBindingHandler) OnUpdate(oldObj, newObj interface{}) {
	if clusterRoleBinding, ok := newObj.(*rbacV1.ClusterRoleBinding); ok {
		this.ClusterRoleBindingMap.Update(clusterRoleBinding)

		// 通知ws客户端
		clusterRoleBindingList := this.ClusterRoleBindingSvc.List()
		page := athena.NewPage(1, 10)
		ws.ClientMap.SendAll(dto.WS{
			Type: "clusterRoleBindings",
			Result: gin.H{
				"ns":   "",
				"data": this.ClusterRoleBindingSvc.Paging(page, clusterRoleBindingList),
			},
		})
	}
}

func (this *ClusterRoleBindingHandler) OnDelete(obj interface{}) {
	if clusterRoleBinding, ok := obj.(*rbacV1.ClusterRoleBinding); ok {
		this.ClusterRoleBindingMap.Delete(clusterRoleBinding)

		// 通知ws客户端
		clusterRoleBindingList := this.ClusterRoleBindingSvc.List()
		page := athena.NewPage(1, 10)
		ws.ClientMap.SendAll(dto.WS{
			Type: "clusterRoleBindings",
			Result: gin.H{
				"ns":   "",
				"data": this.ClusterRoleBindingSvc.Paging(page, clusterRoleBindingList),
			},
		})
	}
}

var _ cache.ResourceEventHandler = &ClusterRoleBindingHandler{}
