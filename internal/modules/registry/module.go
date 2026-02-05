package registry

import (
	"context"

	"github.com/usesnipet/snipet-core-go/internal/infra/provider/file"
	"github.com/usesnipet/snipet-core-go/internal/infra/storage"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewRegistry,
	),
	fx.Invoke(func(
		lc fx.Lifecycle,
		registry *Registry,
		storage *storage.Service,
	) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				registry.Register(file.New(storage))
				return nil
			},
		})
	}),
)
