APP_NAME=snipet-core
DB_NAME=snipet_core_go
DB_USER=postgres
DB_PASS=postgres
DB_HOST=localhost
DB_PORT=5432


ATLAS_ENV=gorm

.PHONY: help run dev build clean db-create migrate migrate-diff migrate-status lint test

help:
	@echo ""
	@echo "Available commands:"
	@echo "  make dev              Run app with air"
	@echo "  make run              Run app"
	@echo "  make build            Build binary"
	@echo "  make clean            Clean build artifacts"
	@echo "  make db-create        Create database if not exists"
	@echo "  make migrate          Apply migrations"
	@echo "  make migrate-diff     Generate new migration"
	@echo "  make migrate-status   Show migration status"
	@echo ""

dev:
	air

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP_NAME) ./cmd/api

clean:
	rm -rf bin tmp

db-create:
	go run ./cmd/tools/db-create

migrate-diff:
	atlas migrate diff $(name) --env $(ATLAS_ENV)
