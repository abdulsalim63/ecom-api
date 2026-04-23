package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/abdulsalim63/ecom-api/internal/adapters/postgresql.sqlc"
	"github.com/abdulsalim63/ecom-api/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting for production scale
	r.Use(middleware.RealIP)    // import for rate limiting, tracing, and analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from GO!!"))
	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/products/{id}", productHandler.FindProductByID)
	r.Post("/product", productHandler.CreateProduct)
	r.Post("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DeleteProduct)
	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("server has started at port %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
