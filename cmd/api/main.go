package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/abdulsalim63/ecom-api/internal/config"
	"github.com/abdulsalim63/ecom-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	srv, err := server.New(ctx, cfg)
	if err != nil {
		slog.Error("failed to create server", "error", err)
		os.Exit(1)
	}

	if err := srv.Run(); err != nil {
		slog.Error("server stopped with error", "error", err)
		os.Exit(1)
	}

	// cfg := config{
	// 	addr: ":8080",
	// 	db: dbConfig{
	// 		dsn: env.GetString("GOOSE_DBSTRING", "host=localhost port=5434 user=postgres password=123123 dbname=ecom sslmode=disable"),
	// 	},
	// }

	// // Logger
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// slog.SetDefault(logger)

	// // Database
	// pool, err := pgxpool.New(ctx, cfg.db.dsn)
	// if err != nil {
	// 	panic(err)
	// }
	// defer pool.Close()

	// slog.Info("Database connected!!", "dsn", cfg.db.dsn)

	// api := application{
	// 	config: cfg,
	// 	db:     pool,
	// }

	// if err := api.run(api.mount()); err != nil {
	// 	slog.Error("server failed to start", "error", err)
	// 	os.Exit(1)
	// }
}
