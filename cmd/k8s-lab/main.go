package main

import (
	"log"
	"net/http"
	"time"

	"github.com/skhanal/k8s-lab/internal/api"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      api.NewRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("server listening on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
