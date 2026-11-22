package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"nimbus/internal/config"
	"nimbus/internal/httpapi"
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

	logger.Info("Starting Nimbus API", "addr", cfg.HTTPAddr)

	err = server.ListenAndServe()

	if err != nil {
		logger.Error("Server error", "error", err)
	}
}
