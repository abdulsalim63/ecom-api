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
	// Rule 1 - Product name must unique
	existedProducts, err := s.repo.ListForAdmin(ctx, product.Filter{
		Search: params.Name,
	})
	if err != nil {
		return 0, err
	}

	if existedProducts != nil {
		return 0, product.ErrExistedProduct
	}

	// Rule 2 - Price and Stock must be greater than 0
	if (params.PriceInCent < 0) || (params.Stock < 0) {
		return 0, product.ErrInvalidPriceOrStock
	}

	createdId, err := s.repo.Create(ctx, params)
	if err != nil {
		return 0, err
	}

	return createdId, nil
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
