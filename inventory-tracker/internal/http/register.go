package http

import (
	"github.com/gin-gonic/gin"
	"github.com/Crstpr/inventory-tracker/internal/http/handler"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health/live", handler.HealthLive)
}