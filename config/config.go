package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config содержит все настройки приложения.
type Config struct {
    ServerAddress string
    PostgresURL   string
    RedisURL      string
}

// LoadConfig загружает конфигурацию из переменных окружения.
func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("Не удалось загрузить файл .env. Используем переменные окружения.")
    }

    return &Config{
        ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
        PostgresURL:   getEnv("POSTGRES_URL", "postgres://user:password@localhost:5432/db?sslmode=disable"),
        RedisURL:      getEnv("REDIS_URL", "localhost:6379"),
    }
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию.
func getEnv(key, defaultValue string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return defaultValue
}
