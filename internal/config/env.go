package config

import (
	"log"
	"os"
	"sync"

	"github.com/go-playground/validator/v10"
)

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

var validate = validator.New()

func validateEnv(env *Env) {
	if err := validate.Struct(env); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Fatalf(
				"❌ Invalid env var: %s (rule: %s)",
				e.Field(),
				e.Tag(),
			)
		}
	}
}

type Env struct {
	// DATABASE
	DB_NAME string `validate:"min=1"`
	DB_USER string `validate:"min=1"`
	DB_PASS string
	DB_HOST string `validate:"hostname|ip"`
	DB_PORT string `validate:"numeric"`
	DB_SSL  string `validate:"oneof=disable require verify-ca verify-full"`

	// REDIS
	REDIS_ADDR string `validate:"hostname_port"`
	REDIS_USER string
	REDIS_PASS string

	// STORAGE
	STORAGE_ENDPOINT   string `validate:"min=1"`
	STORAGE_REGION     string `validate:"min=1"`
	STORAGE_BUCKET     string `validate:"min=1"`
	STORAGE_ACCESS_KEY string `validate:"required,min=1"`
	STORAGE_SECRET_KEY string `validate:"required,min=1"`
	STORAGE_PATH_STYLE bool   `validate:"boolean"`
}

func newEnv() *Env {
	env := &Env{
		DB_NAME: getenv("DB_NAME", "snipet_core_go"),
		DB_USER: getenv("DB_USER", "postgres"),
		DB_PASS: getenv("DB_PASS", "postgres"),
		DB_HOST: getenv("DB_HOST", "localhost"),
		DB_PORT: getenv("DB_PORT", "5432"),
		DB_SSL:  getenv("DB_SSL", "disable"),

		REDIS_ADDR: getenv("REDIS_ADDR", "localhost:6379"),
		REDIS_USER: getenv("REDIS_USER", ""),
		REDIS_PASS: getenv("REDIS_PASS", ""),

		STORAGE_ENDPOINT:   getenv("STORAGE_ENDPOINT", "http://localhost:9000"),
		STORAGE_REGION:     getenv("STORAGE_REGION", "us-east-1"),
		STORAGE_BUCKET:     getenv("STORAGE_BUCKET", "files"),
		STORAGE_ACCESS_KEY: getenv("STORAGE_ACCESS_KEY", "minio"),
		STORAGE_SECRET_KEY: getenv("STORAGE_SECRET_KEY", "adminadmin"),
		STORAGE_PATH_STYLE: getenv("STORAGE_PATH_STYLE", "true") == "true",
	}

	validateEnv(env)
	return env
}

var env *Env
var once sync.Once

func GetEnv() *Env {
	once.Do(func() {
		env = newEnv()
	})
	return env
}
