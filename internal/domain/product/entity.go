package product

import (
	"time"

	filter "github.com/abdulsalim63/ecom-api/internal/domain"
)

// Product is the domain entity.
// It has zero knowledge of sqlc, GORM, HTTP, or any framework.
// This is the single source of truth for what a "product" means in this system.
type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	PriceInCent int32     `json:"price_in_cents"`
	Stock       int32     `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

// -----------------------------------------------------------------
// Command / Query parameter objects
// These are the inputs to use cases — not sqlc types, not HTTP types.
// They belong to the domain layer because they express domain intent.
// -----------------------------------------------------------------

// CreateParams holds the data needed to register a new product.
type CreateParams struct {
	Name        string `json:"name"`
	PriceInCent int32  `json:"price_in_cents"`
	Stock       int32  `json:"stock"`
}

// UpdateParams holds the data allowed to be changed on an existing product.
type UpdateParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PriceInCent int32  `json:"price_in_cents"`
	Stock       int32  `json:"stock"`
}

// Filter holds query parameters for listing products.
type Filter struct {
	PriceMin int32  `json:"price_min"`
	PriceMax int32  `json:"price_max"`
	StockMin int32  `json:"stock_min"`
	StockMax int32  `json:"stock_max"`
	Search   string `json:"search"`
	Page     filter.Pagination
}
