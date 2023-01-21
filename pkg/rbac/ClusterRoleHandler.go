package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	"k8s.io/client-go/tools/cache"
	"pixelk8/src/dto"
	"pixelk8/src/ws"
)

type ClusterRoleHandler struct {
	ClusterRoleMap *ClusterRoleMap     `inject:"-"`
	ClusterRoleSvc *ClusterRoleService `inject:"-"`
}

func (this *ClusterRoleHandler) OnAdd(obj interface{}) {
	if clusterRole, ok := obj.(*rbacV1.ClusterRole); ok {
		this.ClusterRoleMap.Add(clusterRole)

		// 通知ws客户端
		clusterRoleList := this.ClusterRoleSvc.List()
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "clusterRoles",
			Result: gin.H{
				"ns":   "",
				"data": this.ClusterRoleSvc.Paging(page, clusterRoleList),
			},
		})
	}
}

func (this *ClusterRoleHandler) OnUpdate(oldObj, newObj interface{}) {
	if clusterRole, ok := newObj.(*rbacV1.ClusterRole); ok {
		this.ClusterRoleMap.Update(clusterRole)

		// 通知ws客户端
		clusterRoleList := this.ClusterRoleSvc.List()
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "clusterRoles",
			Result: gin.H{
				"ns":   "",
				"data": this.ClusterRoleSvc.Paging(page, clusterRoleList),
			},
		})
	}
}

func (this *ClusterRoleHandler) OnDelete(obj interface{}) {
	if clusterRole, ok := obj.(*rbacV1.ClusterRole); ok {
		this.ClusterRoleMap.Delete(clusterRole)

		// 通知ws客户端
		clusterRoleList := this.ClusterRoleSvc.List()
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "clusterRoles",
			Result: gin.H{
				"ns":   "",
				"data": this.ClusterRoleSvc.Paging(page, clusterRoleList),
			},
		})
	}
}

var _ cache.ResourceEventHandler = &ClusterRoleHandler{}
