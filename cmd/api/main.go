package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gin-gonic/gin"
	"github.com/usesnipet/snipet-core-go/internal/infra/database"
	"github.com/usesnipet/snipet-core-go/internal/infra/http"
	"github.com/usesnipet/snipet-core-go/internal/modules/knowledge"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		database.Module,

		http.Module,

		knowledge.Module,

		fx.Invoke(func(
			engine *gin.Engine,
			controllersIn http.ControllersIn,
		) {
			api := engine.Group("/api")

			for _, controller := range controllersIn.Controllers {
				controller.RegisterRoutes(api)
			}

			engine.Run(":8852")
		}),
	)
	app.Run()
}
