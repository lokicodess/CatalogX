package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lokicodess/CatalogX/internal/domain"
)

type PostgresProductRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresProductRepository(db *pgxpool.Pool) *PostgresProductRepository {
	return &PostgresProductRepository{DB: db}
}

func (p *PostgresProductRepository) Create(ctx context.Context, product *domain.Product) error {

	// TODO: ERROR HANDLING  REMAINING

	stmt := `
INSERT INTO products (name, slug, description, price, sku, stock_quantity, is_active)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at
`
	err := p.DB.QueryRow(ctx, stmt, product.Name, product.Slug, product.Description, product.Price, product.Sku, product.StockQuantity, product.IsActive).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresProductRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {

	// TODO: ERROR HANDLING  REMAINING

	product := &domain.Product{}

	stmt := `
SELECT id, name, slug, description, price, sku, stock_quantity, is_active, created_at, updated_at
FROM products
WHERE id = $1;
`

	err := p.DB.QueryRow(ctx, stmt, id).Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Description,
		&product.Price,
		&product.Sku,
		&product.StockQuantity,
		&product.IsActive,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *PostgresProductRepository) GetAll(ctx context.Context) ([]*domain.Product, error) {
	stmt := `
SELECT id, name, slug, description, price, sku, stock_quantity, is_active, created_at, updated_at
FROM products;
`

	rows, err := p.DB.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*domain.Product{}

	for rows.Next() {
		product := &domain.Product{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.Price,
			&product.Sku,
			&product.StockQuantity,
			&product.IsActive,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
