package main

import (
	"email-classifier/internal/api"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	api.InitClassifier(apiKey)

	http.HandleFunc("/classify", api.ClassifyEmailHandler)
	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
