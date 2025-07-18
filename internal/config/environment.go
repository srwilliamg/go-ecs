package config

import (
	"os"
	"strconv"
)

type EnviromentConfig struct {
	Port     string
	DBHost   string
	DBPort   int
	DBUser   string
	DBPass   string
	DBName   string
	LogLevel string
}

var Envs *EnviromentConfig

func Load() *EnviromentConfig {
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	Envs = &EnviromentConfig{
		Port:     getEnv("PORT", "8080"),
		DBHost:   getEnv("DB_HOST", "localhost"),
		DBPort:   dbPort,
		DBUser:   getEnv("DB_USER", "user"),
		DBPass:   getEnv("DB_PASS", "pass"),
		DBName:   getEnv("DB_NAME", "gopi-db"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
	return Envs
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
