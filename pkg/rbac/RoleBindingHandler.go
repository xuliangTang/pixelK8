package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	rbacV1 "k8s.io/api/rbac/v1"
	"log"
	"pixelk8/src/dto"
	"pixelk8/src/ws"
)

type RoleBindingHandler struct {
	RoleBindingMap *RoleBindingMap     `inject:"-"`
	RoleBindingSvc *RoleBindingService `inject:"-"`
}

func (this *RoleBindingHandler) OnAdd(obj interface{}) {
	roleBinding := obj.(*rbacV1.RoleBinding)
	this.RoleBindingMap.Add(roleBinding)

	// 通知ws客户端
	roleBindingList := this.RoleBindingSvc.ListByNs(roleBinding.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "roleBindings",
		Result: gin.H{
			"ns":   roleBinding.Namespace,
			"data": this.RoleBindingSvc.Paging(page, roleBindingList),
		},
	})
}

func (this *RoleBindingHandler) OnUpdate(oldObj, newObj interface{}) {
	roleBinding := newObj.(*rbacV1.RoleBinding)
	err := this.RoleBindingMap.Update(roleBinding)
	if err != nil {
		log.Println(err)
	} else {
		roleBindingList := this.RoleBindingSvc.ListByNs(roleBinding.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "roleBindings",
			Result: gin.H{
				"ns":   roleBinding.Namespace,
				"data": this.RoleBindingSvc.Paging(page, roleBindingList),
			},
		})
	}
}

func (this *RoleBindingHandler) OnDelete(obj interface{}) {
	roleBinding := obj.(*rbacV1.RoleBinding)
	this.RoleBindingMap.Delete(roleBinding)

	roleBindingList := this.RoleBindingSvc.ListByNs(roleBinding.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "roleBindings",
		Result: gin.H{
			"ns":   roleBinding.Namespace,
			"data": this.RoleBindingSvc.Paging(page, roleBindingList),
		},
	})
}
