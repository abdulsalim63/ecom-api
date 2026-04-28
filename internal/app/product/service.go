package product

import (
	"context"

	"github.com/abdulsalim63/ecom-api/internal/domain/product"
)

type service struct {
	repo product.Repository
}

func NewService(repo product.Repository) product.Service {
	return &service{repo: repo}
}

func (s *service) Add(ctx context.Context, params product.CreateParams) (int64, error) {
	return s.repo.Create(ctx, params)
}

func (s *service) Update(ctx context.Context, params product.UpdateParams) error {
	return nil
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return nil
}

func (s *service) List(ctx context.Context) ([]product.Product, error) {
	return nil, nil
}

func (s *service) Find(ctx context.Context, id int64) (product.Product, error) {
	return product.Product{}, nil
}
