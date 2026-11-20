package main

import (
	"log/slog"
	"net/http"
	"nimbus/internal/httpapi"
	"time"
)

func main() {
	logger := slog.Default()

	router := httpapi.NewRouter()

	server := http.Server{
		Addr:              "localhost:8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("Starting Nimbus API")

	err := server.ListenAndServe()
	
	if err != nil {
		logger.Error("Server error", "error", err)
	}
}
