package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(router *gin.Context)
	Autorization(router *gin.Context)
}
