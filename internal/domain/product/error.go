package product

import "errors"

// Typed domain errors.
// These are defined in the domain layer and used across all layers.
//
// The handler maps these to HTTP status codes.
// The repository translates DB errors (e.g. sql.ErrNoRows) into these.
// The service uses these to express business rule violations.
//
// This means no layer above the repository ever sees a raw DB error.
var (
	// ErrNotFound is returned when a product lookup finds no matching record.
	ErrNotFound = errors.New("product not found")

	// ErrExistedProduct is returned when a product lookup finds existing matching record.
	ErrExistedProduct = errors.New("product with the same name already existed")

	// ErrInvalidPriceOrStock is returned when adding invalid ammount
	ErrInvalidPriceOrStock = errors.New("Price or Stock must be greater than 0")
)
