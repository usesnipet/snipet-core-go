package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbName := getenv("DB_NAME", "snipet_core_go")
	user := getenv("DB_USER", "postgres")
	pass := getenv("DB_PASS", "postgres")
	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "5432")
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		user, pass, host, port,
	)

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
		return
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Database created:", dbName)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
