package api

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	handler := NewHandler()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/api/v1/message", handler.Message)

	return mux
}
