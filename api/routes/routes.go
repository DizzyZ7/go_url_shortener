package routes

import (
	"go_url_shortener/api/handlers"
	"go_url_shortener/database/postgres"
	"go_url_shortener/database/redis"

	"github.com/gin-gonic/gin"
)

// SetupRoutes регистрирует маршруты для приложения.
func SetupRoutes(router *gin.Engine, pgDB *postgres.DB, redisClient *redis.Client) {
	handler := handlers.NewHandler(pgDB, redisClient)

	router.GET("/:shortURL", handler.RedirectURLHandler)
	router.POST("/shorten", handler.ShortenURLHandler)
}
