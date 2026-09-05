package main

import (
	"context"
	"log"
	"os"

	"github.com/Crstpr/inventory-tracker/internal/dbgen"
	httpApp "github.com/Crstpr/inventory-tracker/internal/http"
	"github.com/Crstpr/inventory-tracker/internal/http/handler"
	"github.com/Crstpr/inventory-tracker/internal/repository"
	"github.com/Crstpr/inventory-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("failed to create database pool:", err)
	}
	defer pool.Close()

	queries := dbgen.New(pool)

	productRepository := repository.NewProductRepository(queries)

	productService := service.NewProductService(productRepository)

	productHandler := handler.NewProductHandler(productService)
	router := gin.Default()

	httpApp.RegisterRoutes(router, productHandler)

	if err := router.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}