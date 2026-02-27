package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/subscribe", subscribeHandler)

	http.HandleFunc("/dapr/subscribe", func(w http.ResponseWriter, r *http.Request) {
		subs := []map[string]string{
			{
				"pubsubname": "pubsub",
				"topic": "test",
				"route": "/subscribe",
			},
		}
		json.NewEncoder(w).Encode(subs)
	})

	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	fmt.Println("Received data:", data)
	w.WriteHeader(http.StatusOK)
}