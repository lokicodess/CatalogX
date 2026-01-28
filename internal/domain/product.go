package domain

import "time"

type Product struct {
	ID            string    `json:"id"` // UUID as string
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"` // decimal(10,2)
	Sku           *string   `json:"sku,omitempty"`
	StockQuantity int       `json:"stock_quantity"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
