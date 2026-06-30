package api

import (
	"encoding/json"
	"net/http"
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
	// for the purposes of making the liveness probe fail.
	// c := atomic.AddInt32(&h.count, 1)

	// if c <= h.limit {
	// 	writeJSON(w, http.StatusOK, map[string]string{
	// 		"status": "ok",
	// 	})
	// 	return
	// }

	// writeJSON(w, http.StatusInternalServerError, map[string]string{
	// 	"status": "fail",
	// })

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
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
