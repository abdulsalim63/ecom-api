package product

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int32   `json:"stock"`
}

type CreateParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
	Stock       int32  `json:"stock"`
}

type UpdateParams struct{}
