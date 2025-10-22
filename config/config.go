package config

import (
	"os"
)

// Config содержит все настройки приложения.
type Config struct {
	ServerPort      string
	PostgresDSN     string
	RedisAddress    string
}

// NewConfig загружает конфигурацию из переменных окружения.
func NewConfig() *Config {
	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		PostgresDSN:     getEnv("POSTGRES_DSN", "postgres://user:password@localhost:5432/urlshortener"),
		RedisAddress:    getEnv("REDIS_ADDRESS", "localhost:6379"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
