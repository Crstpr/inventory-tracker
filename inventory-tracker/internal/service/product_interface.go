package service

import (
	"context"

	"github.com/Crstpr/inventory-tracker/internal/domain"
	"github.com/google/uuid"
)

type ProductServiceInterface interface {
	ListProducts(ctx context.Context) ([]domain.Product, error)

	GetProductByID(ctx context.Context, id uuid.UUID) (domain.Product, error)
}