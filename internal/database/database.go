package database

import (
	"database/sql"
	"email-classifier/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./email-classifier.db")
	if err != nil {
		return err
	}

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS classifications (
        id INTEGER NOT NULL PRIMARY KEY,
        subject TEXT,
        body TEXT,
        classification TEXT
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

func SaveClassification(email models.Email, classification string) error {
	stmt, err := db.Prepare("INSERT INTO classifications(subject, body, classification) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(email.Subject, email.Content, classification)
	if err != nil {
		return err
	}
	return nil
}

func GetAllClassifications() ([]models.Email, error) {
	rows, err := db.Query("SELECT id, subject, body, classification FROM classifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classifications []models.Email
	for rows.Next() {
		var email models.Email
		if err := rows.Scan(&email.ID, &email.Subject, &email.Content, &email.Classification); err != nil {
			return nil, err
		}
		classifications = append(classifications, email)
	}
	return classifications, nil
}

func GetClassificationByID(id int) (models.Email, error) {
	var email models.Email
	err := db.QueryRow("SELECT id, subject, body, classification FROM classifications WHERE id = ?", id).Scan(&email.ID, &email.Subject, &email.Content, &email.Classification)
	if err != nil {
		return email, err
	}
	return email, nil
}

func DeleteClassification(id int) error {
	_, err := db.Exec("DELETE FROM classifications WHERE id = ?", id)
	return err
}
