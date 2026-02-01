package database

import (
	"context"
	"database/sql"
	"log"
	"os/exec"

	_ "github.com/lib/pq"
	"github.com/usesnipet/snipet-core-go/internal/config"
	"go.uber.org/fx"
)

type Migrator struct{}

func NewMigrator() *Migrator {
	return &Migrator{}
}
func (m *Migrator) createDatabaseIfNotExists() error {
	dsn := getDsn(true)
	dbName := config.GetEnv().DB_NAME
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM pg_database WHERE datname = $1
		)
	`
	if err := db.QueryRow(query, dbName).Scan(&exists); err != nil {
		log.Fatal(err)
	}

	if exists {
		log.Println("ℹ️ Database already exists:", dbName)
		return nil
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database created:", dbName)
	return nil
}

func (m *Migrator) Run() error {
	log.Println("📦 Running database migrations...")

	if err := m.createDatabaseIfNotExists(); err != nil {
		return err
	}

	dsn := getDsn(false)
	cmd := exec.Command(
		"atlas",
		"migrate",
		"apply",
		"--env",
		"gorm",
		"--url",
		dsn,
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
