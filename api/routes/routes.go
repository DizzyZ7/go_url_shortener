package routes

import (
    "github.com/gin-gonic/gin"
    "go_url_shortener/api/handlers"
    "go_url_shortener/database/postgres"
    "go_url_shortener/database/redis"
    "net/http"
)

// SetupRoutes регистрирует маршруты для приложения.
func SetupRoutes(router *gin.Engine, pgDB *postgres.DB, redisClient *redis.Client) {
    handler := handlers.NewHandler(pgDB, redisClient)

    // Маршрут для отображения веб-интерфейса
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

    // Маршруты API
    router.GET("/:shortURL", handler.RedirectURLHandler)
    router.POST("/shorten", handler.ShortenURLHandler)
}
