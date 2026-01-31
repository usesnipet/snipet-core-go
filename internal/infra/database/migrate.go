package database

import (
	"context"
	"log"
	"os/exec"

	"go.uber.org/fx"
)

type Migrator struct{}

func NewMigrator() *Migrator {
	return &Migrator{}
}

func (m *Migrator) Run() error {
	log.Println("📦 Running database migrations...")

	cmd := exec.Command(
		"atlas",
		"migrate",
		"apply",
		"--env",
		"local",
	)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	return cmd.Run()
}

func RegisterMigration(lc fx.Lifecycle, migrator *Migrator) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return migrator.Run()
		},
	})
}

var MigrationModule = fx.Module(
	"migrations",
	fx.Provide(NewMigrator),
	fx.Invoke(RegisterMigration),
)
