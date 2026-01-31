package database

import (
	"log"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=gin_api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Database connected")
	return db, nil
}

var Module = fx.Module(
	"database",
	fx.Provide(NewGormDB),
)
