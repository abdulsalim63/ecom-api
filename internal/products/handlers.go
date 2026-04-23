package products

import (
	"log"
	"net/http"
	"strconv"

	repo "github.com/abdulsalim63/ecom-api/internal/adapters/postgresql.sqlc"
	"github.com/abdulsalim63/ecom-api/internal/json"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) ListProductById(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var product repo.Product
	for _, p := range products {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if p.ID == id {
			product = p
			break
		}
	}

	if product == (repo.Product{}) {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, product)
}
