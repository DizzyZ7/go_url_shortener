package routes

import "github.com/gin-gonic/gin"
import "go_url_shortener/api/handlers"

// Регистрация маршрутов
func SetupRoutes(router *gin.Engine) {
    router.GET("/:shortURL", handlers.RedirectURLHandler)
    router.POST("/shorten", handlers.ShortenURLHandler)
}
