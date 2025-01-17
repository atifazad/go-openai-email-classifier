package api

import (
	"email-classifier/internal/classifier"
	"email-classifier/internal/models"
	"encoding/json"
	"net/http"
)

var emailClassifier *classifier.Classifier

func InitClassifier(apiKey string) {
	emailClassifier = classifier.NewClassifier(apiKey)
}

func ClassifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	var email models.Email
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := emailClassifier.ClassifyEmail(email.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
