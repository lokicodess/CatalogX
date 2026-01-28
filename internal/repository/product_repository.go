package repository

import (
	"context"

	"github.com/lokicodess/CatalogX/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	GetAll(ctx context.Context) ([]*domain.Product, error)
}
