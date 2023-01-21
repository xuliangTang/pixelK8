package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	coreV1 "k8s.io/api/core/v1"
	"log"
	"pixelk8/src/dto"
	"pixelk8/src/ws"
)

type ServiceAccountHandler struct {
	SaMap             *ServiceAccountMap     `inject:"-"`
	ServiceAccountSvc *ServiceAccountService `inject:"-"`
}

func (this *ServiceAccountHandler) OnAdd(obj interface{}) {
	sa := obj.(*coreV1.ServiceAccount)
	this.SaMap.Add(sa)

	// 通知ws客户端
	saList := this.ServiceAccountSvc.ListByNs(sa.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "serviceAccounts",
		Result: gin.H{
			"ns":   sa.Namespace,
			"data": this.ServiceAccountSvc.Paging(page, saList),
		},
	})
}

func (this *ServiceAccountHandler) OnUpdate(oldObj, newObj interface{}) {
	sa := newObj.(*coreV1.ServiceAccount)
	err := this.SaMap.Update(sa)
	if err != nil {
		log.Println(err)
	} else {
		// 通知ws客户端
		saList := this.ServiceAccountSvc.ListByNs(sa.Namespace)
		page := athena.NewPage(1, 5)
		ws.ClientMap.SendAll(dto.WS{
			Type: "serviceAccounts",
			Result: gin.H{
				"ns":   sa.Namespace,
				"data": this.ServiceAccountSvc.Paging(page, saList),
			},
		})
	}
}

func (this *ServiceAccountHandler) OnDelete(obj interface{}) {
	sa := obj.(*coreV1.ServiceAccount)
	this.SaMap.Delete(sa)

	// 通知ws客户端
	saList := this.ServiceAccountSvc.ListByNs(sa.Namespace)
	page := athena.NewPage(1, 5)
	ws.ClientMap.SendAll(dto.WS{
		Type: "serviceAccounts",
		Result: gin.H{
			"ns":   sa.Namespace,
			"data": this.ServiceAccountSvc.Paging(page, saList),
		},
	})
}
