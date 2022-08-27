package config

import (
	"encoding/json"
	"github.com/fyllekanin/com.monitier.server/task"
	"log"
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
	Services         []*task.Service
}

func getOSEnvironmentOrDefault(env string, def string) string {
	var val = os.Getenv(env)
	if val == "" {
		return def
	} else {
		return val
	}
}

func getServicesConfig() []*task.Service {
	c, err := os.ReadFile("services.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var services []*task.Service
	json.Unmarshal(c, &services)
	return services
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
		Services:         getServicesConfig(),
	}
}
