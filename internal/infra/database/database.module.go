package database

import "go.uber.org/fx"

var Module = fx.Module(
	"database",
	fx.Provide(NewMigrator),
	fx.Provide(NewGormDB),
)
