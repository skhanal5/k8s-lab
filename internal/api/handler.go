package api

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

type Handler struct {
	count int32
	limit int32
}

func NewHandler() *Handler {
	return &Handler{
		limit: 5, // first 5 requests succeed
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	// for the purposes of making the readiness probe fail.
	c := atomic.AddInt32(&h.count, 1)

	if c <= h.limit {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "it is working :)",
		})
		return
	}

	writeJSON(w, http.StatusInternalServerError, map[string]string{
		"status": "it is broken :()",
	})
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "This is a sample API for k8s lab",
		})

	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func (h *Handler) Message(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "hello from go api",
		})

	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	data any,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
