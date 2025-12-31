package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matsumo_and/dapr-work/app/internal/service/user"
	userv1 "github.com/matsumo_and/dapr-work/app/proto/user/v1/userv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()

	// user-serviceのハンドラーを登録
	userHandler := user.NewHandler()
	path, handler := userv1.NewUserServiceHandler(userHandler)
	mux.Handle(path, handler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("user-service listening on %s", addr)

	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
