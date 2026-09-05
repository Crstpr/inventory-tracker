package repository

import (
	"context"

	"github.com/Crstpr/inventory-tracker/internal/dbgen"
	"github.com/Crstpr/inventory-tracker/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type productRepository struct {
	queries *dbgen.Queries
}

func NewProductRepository(
	queries *dbgen.Queries,
) ProductRepositoryInterface {
	return &productRepository{
		queries: queries,
	}
}

func (r *productRepository) ListProducts(
	ctx context.Context,
) ([]domain.Product, error) {

	rows, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]domain.Product, 0, len(rows))

	for _, row := range rows {
		products = append(products, mapListProductRow(row))
	}

	return products, nil
}

func (r *productRepository) GetProductByID(
	ctx context.Context,
	id uuid.UUID,
) (domain.Product, error) {

	productID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	row, err := r.queries.GetProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}

	return mapGetProductByIDRow(row), nil
}

func mapListProductRow(
	row dbgen.ListProductsRow,
) domain.Product {

	var description *string

	if row.Description.Valid {
		description = &row.Description.String
	}

	return domain.Product{
		ID:          uuid.UUID(row.ID.Bytes),
		SKU:         row.Sku,
		Name:        row.Name,
		Description: description,
		Unit:        row.Unit,
		MinStock:    row.MinStock,
		OnHand:      row.OnHand,
		IsActive:    row.IsActive,
		Version:     row.Version,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}

func mapGetProductByIDRow(
	row dbgen.GetProductByIDRow,
) domain.Product {

	var description *string

	if row.Description.Valid {
		description = &row.Description.String
	}

	return domain.Product{
		ID:          uuid.UUID(row.ID.Bytes),
		SKU:         row.Sku,
		Name:        row.Name,
		Description: description,
		Unit:        row.Unit,
		MinStock:    row.MinStock,
		OnHand:      row.OnHand,
		IsActive:    row.IsActive,
		Version:     row.Version,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}
