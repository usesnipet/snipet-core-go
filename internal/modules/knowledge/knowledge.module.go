package knowledge

import (
	"github.com/usesnipet/snipet-core-go/internal/infra/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"knowledge",
	fx.Provide(
		NewRepository,
		NewService,
		fx.Annotate(
			NewController,
			fx.As(new(http.Controller)),
			fx.ResultTags(`group:"controllers"`),
		),
	),
)
