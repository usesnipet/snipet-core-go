package config

import (
	"os"
	"sync"
)

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

type Env struct {
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_SSL  string
}

func newEnv() *Env {
	return &Env{
		DB_NAME: getenv("DB_NAME", "snipet_core_go"),
		DB_USER: getenv("DB_USER", "postgres"),
		DB_PASS: getenv("DB_PASS", "postgres"),
		DB_HOST: getenv("DB_HOST", "localhost"),
		DB_PORT: getenv("DB_PORT", "5432"),
		DB_SSL:  getenv("DB_SSL", "disable"),
	}
}

var env *Env
var once sync.Once

func GetEnv() *Env {
	once.Do(func() {
		env = newEnv()
	})
	return env
}
