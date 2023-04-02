package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"google.golang.org/grpc/codes"
	"log"
)

func main() {
	r := gin.New()

	r.POST("/", func(context *gin.Context) {
		log.Println("post coming")
		req := &v1alpha1.InterceptorRequest{}
		rsp := &v1alpha1.InterceptorResponse{}

		if err := context.ShouldBindJSON(req); err == nil {
			if token, ok := req.Header["X-Jtthink-Token"]; ok && len(token) > 0 && token[0] == "2345" {
				rsp.Status.Code = codes.OK
				rsp.Continue = true
				context.JSON(200, rsp)
				log.Println("validate success")
				return
			} else {
				log.Println(req.Header)
			}
		} else {
			log.Println(err)
		}
		rsp.Status.Code = codes.Unavailable
		rsp.Status.Message = "error param"
		log.Println(503)
		context.JSON(503, rsp)
	})

	r.Run(":80")
}
