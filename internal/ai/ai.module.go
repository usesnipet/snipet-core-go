package ai

import (
	"github.com/usesnipet/snipet-core-go/internal/ai/registry"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"ai",
	fx.Provide(
		registry.NewInMemoryRegistry,
	),
)
