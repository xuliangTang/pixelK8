package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuliangTang/athena/athena"
	"io"
	"net/http"
	"pixelk8/src/properties"
	"pixelk8/src/requests"
	"pixelk8/src/services"
)

// PodCtl @Controller
type PodCtl struct {
	PodService *services.PodService `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (this *PodCtl) pods(ctx *gin.Context) athena.Collection {
	page := athena.NewPageWithCtx(ctx)
	ns := ctx.DefaultQuery("ns", properties.App.K8s.DefaultNs)
	podList := this.PodService.ListByNs(ns)

	return this.PodService.Paging(page, podList)
}

func (this *PodCtl) podAllContainers(ctx *gin.Context) any {
	uri := &requests.PodAllContainersUri{}
	athena.Error(ctx.BindUri(uri))

	return this.PodService.GetPodContainers(uri)
}

func (this *PodCtl) podContainerLogs(ctx *gin.Context) athena.HttpCode {
	uri := &requests.PodContainersLogsUri{}
	query := &requests.PodContainerLogsQuery{}
	athena.Error(ctx.BindUri(uri))
	athena.Error(ctx.BindQuery(query))

	req := this.PodService.GetPodContainerLog(uri, query)

	reader, err := req.Stream(ctx)
	athena.Error(err)
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n > 0 {
			ctx.Writer.Write(buf[0:n])
			ctx.Writer.(http.Flusher).Flush()
		}
	}

	return http.StatusOK
}

func (this *PodCtl) Build(athena *athena.Athena) {
	// 获取pod列表
	athena.Handle("GET", "/pods", this.pods)
	// 获取pod所有容器
	athena.Handle("GET", "/pod/:ns/:pod/containers", this.podAllContainers)
	// 获取pod容器日志
	athena.Handle("GET", "/pod/:ns/:pod/logs", this.podContainerLogs)
}
