package database

import (
	"fmt"

	"github.com/usesnipet/snipet-core-go/internal/config"
)

func getDsn(master bool) string {
	var dbName string
	if master {
		dbName = "postgres"
	} else {
		dbName = config.GetEnv().DB_NAME
	}

	user := config.GetEnv().DB_USER
	pass := config.GetEnv().DB_PASS
	host := config.GetEnv().DB_HOST
	port := config.GetEnv().DB_PORT
	ssl := config.GetEnv().DB_SSL

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, port, dbName, ssl,
	)
}
