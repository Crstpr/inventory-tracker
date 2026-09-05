package repository

import (
	"context"

	"github.com/Crstpr/inventory-tracker/internal/dbgen"
)

type productRepository struct {
	queries *dbgen.Queries
}


func NewProductRepository(queries *dbgen.Queries) ProductRepositoryInterface {
	return &productRepository{
		queries: queries,
	}
}

func (r *productRepository) ListProducts(ctx context.Context) ([]dbgen.ListProductsRow, error) {
	return r.queries.ListProducts(ctx)
}