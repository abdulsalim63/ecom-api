package products

import (
	"context"

	repo "github.com/abdulsalim63/ecom-api/internal/adapters/postgresql.sqlc"
)

type Service interface {
	CreateProduct(ctx context.Context, arg repo.CreateProductParams) (int64, error)
	DeleteProduct(ctx context.Context, id int64) error
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductByID(ctx context.Context, id int64) (repo.Product, error)
	UpdateProduct(ctx context.Context, arg repo.UpdateProductParams) error
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) CreateProduct(ctx context.Context, arg repo.CreateProductParams) (int64, error) {
	return s.repo.CreateProduct(ctx, arg)
}

func (s *svc) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.DeleteProduct(ctx, id)
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) FindProductByID(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *svc) UpdateProduct(ctx context.Context, arg repo.UpdateProductParams) error {
	return s.repo.UpdateProduct(ctx, arg)
}
