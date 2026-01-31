package knowledge

import "github.com/gin-gonic/gin"

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

func (c *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	knowledges := rg.Group("/knowledge")
	knowledges.GET("", c.list)
}

func (c *Controller) list(ctx *gin.Context) {
	ctx.JSON(200, []string{"a", "b", "c"})
}
