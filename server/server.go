package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Create the mux router
func Router(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	return r
}

// Create the server
func CreateServer(ctx context.Context, port string) *http.Server {
	router := Router(ctx)

	return &http.Server{
		Addr:              port,
		Handler:           router,
		IdleTimeout:       time.Second * 60,
		WriteTimeout:      time.Second * 120,
		ReadTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 120,
	}
}

// Initialize the server on port 8080
