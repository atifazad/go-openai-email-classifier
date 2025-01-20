package api

import (
	"email-classifier/internal/classifier"
	"email-classifier/internal/database"
	"email-classifier/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
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

	err = database.SaveClassification(email, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllClassificationsHandler(w http.ResponseWriter, r *http.Request) {
	classifications, err := database.GetAllClassifications()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(classifications)
}

func GetClassificationByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	classification, err := database.GetClassificationByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(classification)
}

func DeleteClassificationHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteClassification(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
