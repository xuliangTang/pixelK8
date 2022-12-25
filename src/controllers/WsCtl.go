package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"log"
	"net/http"
	"pixelk8/src/ws"
)

// WsCtl @Controller
type WsCtl struct{}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (*WsCtl) connect(ctx *gin.Context) athena.HttpCode {
	client, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}

	ws.ClientMap.Store(client)
	return http.StatusOK
}

func (this *WsCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/ws", this.connect)
}
