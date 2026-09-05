package service

import (
	"context"
	"github.com/Crstpr/inventory-tracker/internal/dbgen"
	"github.com/Crstpr/inventory-tracker/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type productUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewProductService(productRepository repository.ProductRepositoryInterface) ProductServiceInterface {
	return &productUseCase{
		productRepository: productRepository,
	}
}

func (s *productUseCase) ListProducts(ctx context.Context) ([]dbgen.ListProductsRow, error) {
	products, err := s.productRepository.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	if products == nil {
		return []dbgen.ListProductsRow{}, nil
	}
	return products, nil
}

func (s *productUseCase) GetProductByID(ctx context.Context, id pgtype.UUID) (dbgen.GetProductByIDRow, error) {
	product, err := s.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return dbgen.GetProductByIDRow{}, err
	}
	return product, nil
}