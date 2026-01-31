package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoutes(group *gin.RouterGroup)
}

type ControllersIn struct {
	fx.In

	Controllers []Controller `group:"controllers"`
}
