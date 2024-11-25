package settings

import "os"

type Settings struct {
	Port string
}

var instance *Settings

func InitializeFromEnv() *Settings {
	instance := &Settings{
		Port: getEnv("APP_PORT", "8080"),
	}

	return instance
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
