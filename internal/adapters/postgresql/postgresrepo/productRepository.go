package postgresrepo

import (
	"context"

	sqlcgen "github.com/abdulsalim63/ecom-api/internal/adapters/postgresql/generated"
	"github.com/abdulsalim63/ecom-api/internal/domain/product"
)

type productRepository struct {
	q sqlcgen.Querier
}

func NewProductRepository(q sqlcgen.Querier) product.Repository {
	return &productRepository{q: q}
}

func (r *productRepository) Create(ctx context.Context, params product.CreateParams) (int64, error) {
	return r.q.CreateProduct(ctx, sqlcgen.CreateProductParams{
		Name:         params.Name,
		PriceInCents: params.Price,
		Stock:        params.Stock,
	})
}

func (r *productRepository) Update(ctx context.Context, params product.UpdateParams) error {
	return nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	return nil
}

func (r *productRepository) List(ctx context.Context) ([]product.Product, error) {
	return nil, nil
}

func (r *productRepository) Find(ctx context.Context, id int64) (product.Product, error) {
	return product.Product{}, nil
}
