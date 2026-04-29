package product

import "context"

type Service interface {
	Add(ctx context.Context, params CreateParams) (int64, error)
	Update(ctx context.Context, params UpdateParams) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]Product, error) // will implement filter and pagination
	Find(ctx context.Context, id int64) (Product, error)
}
