package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	"log"
	"pixelk8/src/dto"
	"pixelk8/src/ws"
)

type RoleHandler struct {
	RoleMap *RoleMap     `inject:"-"`
	RoleSvc *RoleService `inject:"-"`
}

func (this *RoleHandler) OnAdd(obj interface{}) {
	role := obj.(*rbacV1.Role)
	this.RoleMap.Add(role)

	// 通知ws客户端
	roleList := this.RoleSvc.ListByNs(role.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "roles",
		Result: gin.H{
			"ns":   role.Namespace,
			"data": this.RoleSvc.Paging(page, roleList),
		},
	})
}

func (this *RoleHandler) OnUpdate(oldObj, newObj interface{}) {
	role := newObj.(*rbacV1.Role)
	err := this.RoleMap.Update(role)
	if err != nil {
		log.Println(err)
	} else {
		roleList := this.RoleSvc.ListByNs(role.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "roles",
			Result: gin.H{
				"ns":   role.Namespace,
				"data": this.RoleSvc.Paging(page, roleList),
			},
		})
	}
}

func (this *RoleHandler) OnDelete(obj interface{}) {
	role := obj.(*rbacV1.Role)
	this.RoleMap.Delete(role)

	roleList := this.RoleSvc.ListByNs(role.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "roles",
		Result: gin.H{
			"ns":   role.Namespace,
			"data": this.RoleSvc.Paging(page, roleList),
		},
	})
}
