package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matsumo_and/dapr-work/app/auth-service/internal/handler"
	"github.com/matsumo_and/dapr-work/app/auth-service/proto/authv1/authv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	authHandler := handler.NewAuthHandler()

	mux := http.NewServeMux()
	path, handler := authv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(path, handler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := ":8080"
	fmt.Printf("auth-service listening on %s\n", addr)

	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
