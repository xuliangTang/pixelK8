package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
	"pixelk8/src/requests"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

// WsCtl @Controller
type WsCtl struct {
	PodService *services.PodService `inject:"-"`
}

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

func (this *WsCtl) podContainerTerminal(ctx *gin.Context) athena.HttpCode {
	uri := &requests.PodContainerTerminalUri{}
	athena.Error(ctx.BindUri(uri))

	query := &requests.PodContainerTerminalQuery{}
	athena.Error(ctx.BindQuery(query))

	wsClient, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return http.StatusBadRequest
	}

	shellClient := ws.NewWsShellClient(wsClient)
	exec, err := this.PodService.HandlerCommand(uri, query)
	if err != nil {
		return http.StatusBadRequest
	}

	err = exec.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:  shellClient,
		Stdout: shellClient,
		Stderr: shellClient,
		Tty:    true,
	})

	return http.StatusOK
}

func (this *WsCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/ws", this.connect)
	// 连接pod容器终端
	athena.Handle("GET", "/ws/pod/:ns/:pod/terminal", this.podContainerTerminal)
}
