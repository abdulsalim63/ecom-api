package product

import (
	"net/http"

	"github.com/abdulsalim63/ecom-api/internal/domain/product"
	"github.com/abdulsalim63/ecom-api/pkg/validator"
)

type Handler struct {
	svc       product.Service
	validator *validator.Validator
}

func New(svc product.Service) *Handler {
	return &Handler{svc: svc, validator: validator.New()}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
