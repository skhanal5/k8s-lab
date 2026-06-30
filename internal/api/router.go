package api

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	handler := NewHandler()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/ready", handler.Ready)
	mux.HandleFunc("/memory", handler.Memory)
	mux.HandleFunc("/api/v1/message", handler.Message)

	return mux
}
