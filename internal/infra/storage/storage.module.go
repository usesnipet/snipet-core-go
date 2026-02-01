package storage

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewS3Client,
		NewStorageService,
	),
)
