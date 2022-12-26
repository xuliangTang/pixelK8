package dto

import (
	"github.com/gin-gonic/gin"
)

type WS struct {
	Type   string
	Result gin.H
}
