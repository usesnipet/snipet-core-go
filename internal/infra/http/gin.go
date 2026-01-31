package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGin() *gin.Engine {
	return gin.Default()
}

var Module = fx.Module(
	"http",
	fx.Provide(NewGin),
)
