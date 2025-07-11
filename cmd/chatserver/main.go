// It implements the main entry point for the chat server application.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the chat server
	fmt.Println("Starting chat server...")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})

	fmt.Println("Running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
