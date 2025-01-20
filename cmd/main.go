package main

import (
	"email-classifier/internal/api"
	"email-classifier/internal/database"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	api.InitClassifier(apiKey)

	http.HandleFunc("/classify", api.ClassifyEmailHandler)
	http.HandleFunc("/classifications", api.GetAllClassificationsHandler)
	http.HandleFunc("/classification", api.GetClassificationByIDHandler)
	http.HandleFunc("/delete-classification", api.DeleteClassificationHandler)
	log.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
