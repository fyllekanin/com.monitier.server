package config

import (
	"os"
)

type AppConfig struct {
	Port             string
	DatabaseType     string
	DatabaseHost     string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     string
}

func getOSEnvironmentOrDefault(env string, def string) string {
	var val = os.Getenv(env)
	if val == "" {
		return def
	} else {
		return val
	}
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Port:             getOSEnvironmentOrDefault("SERVICE_PORT", "8080"),
		DatabaseType:     getOSEnvironmentOrDefault("DATABASE_TYPE", "SQLite"),
		DatabaseHost:     getOSEnvironmentOrDefault("DATABASE_HOST", "localhost"),
		DatabaseName:     getOSEnvironmentOrDefault("DATABASE_NAME", "monitier"),
		DatabaseUsername: getOSEnvironmentOrDefault("DATABASE_USERNAME", "username"),
		DatabasePassword: getOSEnvironmentOrDefault("DATABASE_PASSWORD", "password"),
		DatabasePort:     getOSEnvironmentOrDefault("DATABASE_PORT", ""),
	}
}
