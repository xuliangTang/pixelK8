package dto

import (
	"github.com/gin-gonic/gin"
)

type WS struct {
	Type   string `json:"type"`
	Result gin.H  `json:"result"`
}
