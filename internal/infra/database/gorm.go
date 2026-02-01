package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(migrator *Migrator) (*gorm.DB, error) {
	migrator.Run()
	db, err := gorm.Open(postgres.Open(getDsn(false)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Database connected")
	return db, nil
}
