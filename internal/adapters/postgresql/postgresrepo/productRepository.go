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

func toDomain(row sqlcgen.Product) *product.Product {
	return &product.Product{
		ID:          row.ID,
		Name:        row.Name,
		PriceInCent: row.PriceInCents,
		Stock:       row.Stock,
		CreatedAt:   row.CreatedAt.Time,
	}
}

func (r *productRepository) Create(ctx context.Context, params product.CreateParams) (int64, error) {
	row, err := r.q.CreateProduct(ctx, sqlcgen.CreateProductParams{
		Name:         params.Name,
		PriceInCents: params.PriceInCent,
		Stock:        params.Stock,
	})
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (r *productRepository) Update(ctx context.Context, params product.UpdateParams) error {
	err := r.q.UpdateProduct(ctx, sqlcgen.UpdateProductParams{
		Name:         params.Name,
		PriceInCents: params.PriceInCent,
		Stock:        params.Stock,
		ID:           params.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) error {
	err := r.q.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) ListForCustomer(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	data, err := r.q.ListProductsCustomer(ctx, sqlcgen.ListProductsCustomerParams{
		PriceInCents:   filter.PriceMin,
		PriceInCents_2: filter.PriceMax,
		Search:         filter.Search,
		Limit:          filter.Page.Limit,
		Offset:         filter.Page.Offset,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]product.Product, len(data))
	for i, d := range data {
		resp[i] = *toDomain(d)
	}
	return resp, nil
}

func (r *productRepository) ListForAdmin(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	data, err := r.q.ListProductsAdmin(ctx, sqlcgen.ListProductsAdminParams{
		PriceInCents:   filter.PriceMin,
		PriceInCents_2: filter.PriceMax,
		Stock:          filter.StockMin,
		Stock_2:        filter.StockMax,
		Search:         filter.Search,
		Limit:          filter.Page.Limit,
		Offset:         filter.Page.Offset,
	})
	if err != nil {
		return nil, err
	}

	resp := make([]product.Product, len(data))
	for i, d := range data {
		resp[i] = *toDomain(d)
	}
	return resp, nil
}

func (r *productRepository) Find(ctx context.Context, id int64) (product.Product, error) {
	data, err := r.q.FindProductByID(ctx, id)
	if err != nil {
		return product.Product{}, err
	}

	return *toDomain(data), nil
}
