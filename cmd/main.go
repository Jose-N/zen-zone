package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

func main() {
	mux := http.NewServeMux()

	// nux.HandleFunc("/", helloHandler)
	mux.Handle("GET /", helloHandler())
	slog.Info("")
	http.ListenAndServe("localhost:8080", mux)
}
