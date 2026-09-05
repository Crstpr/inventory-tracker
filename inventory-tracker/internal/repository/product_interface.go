package repository

import (
	"context"
	"github.com/Crstpr/inventory-tracker/internal/dbgen"
)

type ProductRepositoryInterface interface {
	ListProducts(ctx context.Context) ([]dbgen.ListProductsRow, error)
}