package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matsumo_and/dapr-work/app/user-service/internal/handler"
	"github.com/matsumo_and/dapr-work/app/user-service/proto/userv1/userv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	userHandler := handler.NewUserHandler()

	mux := http.NewServeMux()
	path, handler := userv1connect.NewUserServiceHandler(userHandler)
	mux.Handle(path, handler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := ":8081"
	fmt.Printf("user-service listening on %s\n", addr)

	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
