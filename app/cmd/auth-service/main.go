package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matsumo_and/dapr-work/app/internal/service/auth"
	"github.com/matsumo_and/dapr-work/app/proto/auth/v1/authv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// auth-serviceのハンドラーを登録
	authHandler := auth.NewHandler()
	path, handler := authv1connect.NewAuthServiceHandler(authHandler)
	mux.Handle(path, handler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("auth-service listening on %s", addr)

	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
