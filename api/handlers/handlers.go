package handlers

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go_url_shortener/database/postgres"
	"go_url_shortener/database/redis"
	"go_url_shortener/models"
)

// Handler содержит зависимости для обработчиков.
type Handler struct {
	pgDB        *postgres.DB
	redisClient *redis.Client
}

// NewHandler создаёт и возвращает новый Handler.
func NewHandler(pgDB *postgres.DB, redisClient *redis.Client) *Handler {
	return &Handler{pgDB: pgDB, redisClient: redisClient}
}

// ShortenURLHandler обрабатывает запросы на сокращение URL.
func (h *Handler) ShortenURLHandler(c *gin.Context) {
	var input models.URL
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат данных"})
		return
	}

	// Проверка, что URL не пуст
	if input.OriginalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL не может быть пустым"})
		return
	}

	// Генерация уникального короткого кода
	shortURL := generateShortURL()

	// Сохранение в PostgreSQL
	query := "INSERT INTO urls (short_url, original_url) VALUES ($1, $2) RETURNING id, created_at"
	err := h.pgDB.QueryRow(query, shortURL, input.OriginalURL).Scan(&input.ID, &input.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить URL"})
		return
	}

	// Кэширование в Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = h.redisClient.Set(ctx, shortURL, input.OriginalURL, 24*time.Hour).Err() // Кэш на 24 часа
	if err != nil {
		fmt.Printf("Не удалось закэшировать в Redis: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": fmt.Sprintf("%s/%s", c.Request.Host, shortURL),
	})
}

// generateShortURL генерирует уникальный короткий код.
func generateShortURL() string {
	b := make([]byte, 4) // Генерируем 8-символьную строку
	if _, err := rand.Read(b); err != nil {
		return "" // В случае ошибки
	}
	return hex.EncodeToString(b)
}

// RedirectURLHandler обрабатывает запросы на перенаправление.
func (h *Handler) RedirectURLHandler(c *gin.Context) {
	shortURL := c.Param("shortURL")

	// Поиск в кэше Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	originalURL, err := h.redisClient.Get(ctx, shortURL).Result()

	if err == nil {
		// Перенаправление, если найдено в Redis
		c.Redirect(http.StatusMovedPermanently, originalURL)
		return
	}

	// Поиск в PostgreSQL, если не найдено в Redis
	var urlEntry models.URL
	query := "SELECT original_url FROM urls WHERE short_url = $1"
	err = h.pgDB.QueryRow(query, shortURL).Scan(&urlEntry.OriginalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ссылка не найдена"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка базы данных"})
		return
	}

	// Кэширование результата в Redis на случай повторного запроса
	err = h.redisClient.Set(ctx, shortURL, urlEntry.OriginalURL, 24*time.Hour).Err()
	if err != nil {
		fmt.Printf("Не удалось закэшировать в Redis: %v\n", err)
	}

	// Перенаправление
	c.Redirect(http.StatusMovedPermanently, urlEntry.OriginalURL)
}
