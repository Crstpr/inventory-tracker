package http

import (
	"github.com/Crstpr/inventory-tracker/internal/http/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	productHandler *handler.ProductHandler,
) {
	router.GET("/health/live", handler.HealthLive)

	api := router.Group("/api/v1")

	api.GET("/products", productHandler.ListProducts)
	api.GET("/products/:id", productHandler.GetProductByID)
}