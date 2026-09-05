package service

import (
	"context"

	"github.com/Crstpr/inventory-tracker/internal/domain"
	"github.com/Crstpr/inventory-tracker/internal/repository"
	"github.com/google/uuid"
)

type productService struct {
	productRepository repository.ProductRepositoryInterface
}

func NewProductService(productRepository repository.ProductRepositoryInterface) ProductServiceInterface {
	return &productService{
		productRepository: productRepository,
	}
}

func (u *productService) ListProducts(ctx context.Context) ([]domain.Product, error) {

	products, err := u.productRepository.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *productService) GetProductByID(ctx context.Context, id uuid.UUID) (domain.Product, error) {

	product, err := u.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}