package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Jose-N/zen-zone/internal/http/indexhandler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	mux.Handle("GET /", indexhandler.IndexHandler())

	logger.Info("Starting server", "port", 8080)
	http.ListenAndServe("localhost:8080", mux)
}
