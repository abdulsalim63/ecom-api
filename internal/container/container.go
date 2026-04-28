package container

import (
	"net/http"

	sqlcgen "github.com/abdulsalim63/ecom-api/internal/adapters/postgresql/generated"
	postgresrepo "github.com/abdulsalim63/ecom-api/internal/adapters/postgresql/postgresrepo"
	productapp "github.com/abdulsalim63/ecom-api/internal/app/product"
	producthandler "github.com/abdulsalim63/ecom-api/internal/handler/product"
	"github.com/abdulsalim63/ecom-api/internal/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	Router http.Handler
}

func Build(pool *pgxpool.Pool) *Container {
	// Infrastructure — sqlc only exists in this block
	queries := sqlcgen.New(pool)

	// Repositories — implement domain interfaces, hide sqlc
	productRepo := postgresrepo.NewProductRepository(queries)
	// userRepo    := postgresrepo.NewUserRepository(queries)
	// orderRepo   := postgresrepo.NewOrderRepository(queries)

	// Application services — use domain interfaces, no sqlc
	productSvc := productapp.NewService(productRepo)
	// userSvc    := userapp.NewService(userRepo)
	// orderSvc   := orderapp.NewService(orderRepo)

	// Handlers — use application service interfaces only
	productHandler := producthandler.New(productSvc)
	// userHandler    := userhandler.New(userSvc)
	// orderHandler   := orderhandler.New(orderSvc)

	// Router — mounts handlers, no business logic
	r := router.New(productHandler)

	return &Container{Router: r}
}
