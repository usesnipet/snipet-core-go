package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/usesnipet/snipet-core-go/internal/infra/database"
	http_server "github.com/usesnipet/snipet-core-go/internal/infra/http-server"
	"github.com/usesnipet/snipet-core-go/internal/infra/queue"
	"github.com/usesnipet/snipet-core-go/internal/infra/storage"
	knowledge_source "github.com/usesnipet/snipet-core-go/internal/modules/knowledge-source"
	"github.com/usesnipet/snipet-core-go/internal/modules/registry"
	"github.com/usesnipet/snipet-core-go/internal/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// INFRA
		database.Module,
		queue.Module,
		storage.Module,

		service.Module,

		// MODULES
		knowledge_source.Module,
		registry.Module,

		// HTTP
		http_server.Module,
	)
	app.Run()
}
