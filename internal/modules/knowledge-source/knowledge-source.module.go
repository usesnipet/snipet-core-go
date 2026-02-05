package knowledge_source

import (
	"github.com/go-playground/validator/v10"
	http_server "github.com/usesnipet/snipet-core-go/internal/infra/http-server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"knowledge_source",
	fx.Provide(
		NewRepository,
		NewService,
		func() *validator.Validate { return validator.New() },
		fx.Annotate(
			NewController,
			fx.As(new(http_server.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
)
