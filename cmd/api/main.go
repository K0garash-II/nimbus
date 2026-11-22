package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"nimbus/internal/config"
	"nimbus/internal/httpapi"
	"nimbus/internal/storage/postgres"
)

func main() {
	logger := slog.Default()

	router := httpapi.NewRouter()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Config error", "error", err)
		os.Exit(1)
	}

	server := http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, canсel := context.WithTimeout(context.Background(), 5*time.Second)

	store, err := postgres.NewStore(ctx, cfg.DatabaseURL)

	canсel()

	if err != nil {
		logger.Error("DB error", "error", err)
		os.Exit(1)
	}

	defer store.Close()

	logger.Info("Connected to PostgreSQL")
	logger.Info("Starting Nimbus API", "addr", cfg.HTTPAddr)

	err = server.ListenAndServe()

	if err != nil {
		logger.Error("Server error", "error", err)
	}
}
