package main

import (
	"github.com/gin-gonic/gin"

	httpApp "github.com/Crstpr/inventory-tracker/internal/http"
)

func main() {
	router := gin.Default()

	httpApp.RegisterRoutes(router)

	router.Run("127.0.0.1:8080")
}