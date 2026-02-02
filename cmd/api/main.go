package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/usesnipet/snipet-core-go/internal/infra/database"
	"github.com/usesnipet/snipet-core-go/internal/infra/http"
	"github.com/usesnipet/snipet-core-go/internal/infra/queue"
	"github.com/usesnipet/snipet-core-go/internal/modules/knowledge"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// INFRA
		database.Module,
		queue.Module,

		// MODULES
		knowledge.Module,

		// HTTP
		http.Module,
	)
	app.Run()
}
