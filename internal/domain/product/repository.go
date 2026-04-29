package product

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, params CreateParams) (int64, error)
	Update(ctx context.Context, params UpdateParams) error
	Delete(ctx context.Context, id int64) error
	ListForCustomer(ctx context.Context, filter Filter) ([]Product, error)
	ListForAdmin(ctx context.Context, filter Filter) ([]Product, error)
	Find(ctx context.Context, id int64) (Product, error)
}
