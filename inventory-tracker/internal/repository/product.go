package repository

import (
	"context"

	"github.com/Crstpr/inventory-tracker/internal/dbgen"
	"github.com/jackc/pgx/v5/pgtype"
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

func (r *productRepository) GetProductByID(ctx context.Context, id pgtype.UUID) (dbgen.GetProductByIDRow, error) {
	return r.queries.GetProductByID(ctx, id)
}
