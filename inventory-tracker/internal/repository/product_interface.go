package repository

import (
	"context"

	"github.com/Crstpr/inventory-tracker/internal/dbgen"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProductRepositoryInterface interface {
	ListProducts(ctx context.Context) ([]dbgen.ListProductsRow, error)
	GetProductByID(ctx context.Context, id pgtype.UUID) (dbgen.GetProductByIDRow, error)
}