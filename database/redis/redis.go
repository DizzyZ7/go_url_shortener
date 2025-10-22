package redis

import (
    "context"
    "log"
    "time"

    "github.com/redis/go-redis/v9"
)

// Client представляет клиент для работы с кэшем Redis.
type Client struct {
    *redis.Client
}

// NewClient устанавливает соединение с Redis и возвращает клиент.
func NewClient(redisURL string) *Client {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    rdb := redis.NewClient(&redis.Options{
        Addr: redisURL,
    })

    if _, err := rdb.Ping(ctx).Result(); err != nil {
        log.Fatalf("Не удалось подключиться к Redis: %v", err)
    }
    log.Println("Успешное подключение к Redis!")
    return &Client{rdb}
}
