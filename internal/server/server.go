package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abdulsalim63/ecom-api/internal/config"
	"github.com/abdulsalim63/ecom-api/internal/container"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg  *config.Config
	http *http.Server
}

func New(ctx context.Context, cfg *config.Config) (*Server, error) {
	// Connection pool — not a single connection
	pool, err := pgxpool.New(ctx, cfg.Database.DSN())
	if err != nil {
		return nil, fmt.Errorf("server: connect database: %w", err)
	}

	// Verify the connection is actually alive
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("server: ping database: %w", err)
	}
	slog.Info("database connected", "host", cfg.Database.Host, "db", cfg.Database.Name)

	// Wire all dependencies — no sqlc imports here
	c := container.Build(pool)

	httpServer := &http.Server{
		Addr:         cfg.App.Addr,
		Handler:      c.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{cfg: cfg, http: httpServer}, nil
}

func (s *Server) Run() error {
	// Start server in background
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("server starting", "addr", s.http.Addr)
		if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	// Wait for shutdown signal or server error
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case sig := <-quit:
		slog.Info("shutdown signal received", "signal", sig)
	}

	// Graceful shutdown — give in-flight requests 30 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("server: graceful shutdown failed: %w", err)
	}

	slog.Info("server stopped gracefully")
	return nil
}
