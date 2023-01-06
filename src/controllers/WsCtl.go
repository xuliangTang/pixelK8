package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"golang.org/x/crypto/ssh"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
	"pixelk8/src/ws"
)

// WsCtl @Controller
type WsCtl struct {
	PodService  *services.PodService  `inject:"-"`
	NodeService *services.NodeService `inject:"-"`
}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (*WsCtl) connect(ctx *gin.Context) (v athena.Void) {
	client, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}

	ws.ClientMap.Store(client)

	return
}

func (this *WsCtl) podContainerTerminal(ctx *gin.Context) (v athena.Void) {
	uri := &requests.PodContainerTerminalUri{}
	athena.Error(ctx.BindUri(uri))

	query := &requests.PodContainerTerminalQuery{}
	athena.Error(ctx.BindQuery(query))

	wsClient, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	shellClient := ws.NewWsShellClient(wsClient)
	exec, err := this.PodService.HandlerCommand(uri, query)
	if err != nil {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	err = exec.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:  shellClient,
		Stdout: shellClient,
		Stderr: shellClient,
		Tty:    true,
	})

	return
}

func (this *WsCtl) nodeTerminal(ctx *gin.Context) (v athena.Void) {
	uri := &requests.NodeTerminalUri{}
	athena.Error(ctx.BindUri(uri))

	wsClient, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	nodeInfo, ok := properties.App.K8s.Nodes[uri.Name]
	if !ok {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	shellClient := ws.NewWsShellClient(wsClient)
	session, err := this.NodeService.SSHConnect(nodeInfo.Username, nodeInfo.Password, nodeInfo.Host, nodeInfo.Port)
	if err != nil {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	defer session.Close()
	session.Stdout = shellClient
	session.Stderr = shellClient
	session.Stdin = shellClient

	var nodeShellModes = ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm-256color", 300, 500, nodeShellModes)
	if err != nil {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	err = session.Run("sh")
	if err != nil {
		ctx.Set(athena.CtxHttpStatusCode, http.StatusBadRequest)
		return
	}

	return
}

func (this *WsCtl) Build(athena *athena.Athena) {
	athena.Handle("GET", "/ws", this.connect)
	// 连接pod容器终端
	athena.Handle("GET", "/ws/pod/:ns/:pod/terminal", this.podContainerTerminal)
	// 连接node终端
	athena.Handle("GET", "/ws/node/:node/terminal", this.nodeTerminal)
}
