package redis

import (
    "context"
    "log"

    "github.com/redis/go-redis/v9"
)

// Client содержит клиент Redis.
var Client *redis.Client

// InitRedis устанавливает соединение с Redis.
func InitRedis(redisURL string) {
    Client = redis.NewClient(&redis.Options{
        Addr: redisURL,
    })

    // Проверяем соединение
    ctx := context.Background()
    _, err := Client.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Не удалось подключиться к Redis: %v", err)
    }

    log.Println("Успешно подключились к Redis!")
}
