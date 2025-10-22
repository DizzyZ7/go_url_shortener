package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "go_url_shortener/api/routes"
    "go_url_shortener/config"
    "go_url_shortener/database/postgres"
    "go_url_shortener/database/redis"
)

func main() {
    // Загрузка конфигурации
    cfg := config.LoadConfig()

    // Подключение к базам данных
    pgDB := postgres.NewDB(cfg.PostgresURL)
    defer pgDB.Close()
    pgDB.Migrate()

    redisClient := redis.NewClient(cfg.RedisURL)
    defer redisClient.Close()

    // Настройка Gin-роутера
    router := gin.Default()
    router.LoadHTMLGlob("views/*") // Загрузка HTML-шаблонов

    routes.SetupRoutes(router, pgDB, redisClient)

    // Запуск сервера
    log.Printf("Сервер запущен на %s", cfg.ServerAddress)
    if err := router.Run(cfg.ServerAddress); err != nil {
        log.Fatalf("Не удалось запустить сервер: %v", err)
    }
}
