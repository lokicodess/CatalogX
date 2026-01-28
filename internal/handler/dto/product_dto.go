package dto

type CreateProductRequest struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	SKU           *string `json:"sku,omitempty"`
	CategoryID    int     `json:"category_id"` // keep this even if not in DB yet
	StockQuantity int     `json:"stock_quantity"`
}
