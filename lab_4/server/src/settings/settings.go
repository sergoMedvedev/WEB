package settings

import (
	"os"
)

type Settings struct {
	Port             string
	PortPostgres     string
	HostPostgres     string
	UserPostgres     string
	PasswordPostgres string
	DbNamePostgres   string
	UrlPostgres      string
}

var StrForConnect string

var instance *Settings

func InitializeFromEnv() {
	instance = &Settings{
		Port:             getEnv("APP_PORT", "8080"),
		PortPostgres:     getEnv("PORT_POSTGRES", "5432"),
		HostPostgres:     getEnv("HOST_POSTGRES", "127.0.0.1"),
		UserPostgres:     getEnv("USER_POSTGRES", "postgres"),
		PasswordPostgres: getEnv("PASSWORD_POSTGRES", "password"),
		DbNamePostgres:   getEnv("DB_NAME_POSTGRES", "test"),
		UrlPostgres:      getEnv("URL_POSTGRES", "postgres://postgres:password@127.0.0.1:5432/test?sslmode=disable"),
	}
}

func Instance() *Settings {
	return instance
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
