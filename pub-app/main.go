package main

import (
	"log"
	"net/http"

	"github.com/dapr/go-sdk/client"
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/publish", publishHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := client.NewClient()
	if err != nil {
		log.Fatalf("Failed to create Dapr client: %v", err)
	}

	err = client.PublishEvent(ctx, "pubsub", "test", []byte("Hello, World from pub-app!"))
	if err != nil {
		log.Fatalf("Failed to publish event: %v", err)
	}
}