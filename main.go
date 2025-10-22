package main

import (
    "github.com/gin-gonic/gin"
    "go_url_shortener/api/routes"
)

func main() {
    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8080")
}
