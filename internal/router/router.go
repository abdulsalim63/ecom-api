package router

import (
	"net/http"
	"time"

	producthandler "github.com/abdulsalim63/ecom-api/internal/handler/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(
	product *producthandler.Handler,
	// user    *userhandler.Handler,
	// order   *orderhandler.Handler,
) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// API v1
	r.Route("/api/v1", func(r chi.Router) {
		// Products — consistent plural, correct HTTP methods
		r.Route("/products", func(r chi.Router) {
			r.Get("/", product.ListProducts)
			// r.Post("/", product.AddProduct)
			// r.Get("/{id}", product.FindProductByID)
			// r.Put("/{id}", product.UpdateProduct) // PUT not POST
			// r.Delete("/{id}", product.DeleteProduct)
		})

		// r.Route("/users", func(r chi.Router) { ... })
		// r.Route("/orders", func(r chi.Router) { ... })
	})

	return r
}
