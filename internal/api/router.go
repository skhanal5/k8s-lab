package api

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	handler := NewHandler()

	mux.HandleFunc("/healthz", handler.Health)
	mux.HandleFunc("/api/v1/message", handler.Message)

	return mux
}
