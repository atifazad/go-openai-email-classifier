package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase(dataSourceName string) {
    var err error
    db, err = sql.Open("sqlite3", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    createTableQuery := `CREATE TABLE IF NOT EXISTS classifications (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email_content TEXT,
        classification_result TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatal(err)
    }
}

func StoreClassification(emailContent string, classificationResult string) error {
    insertQuery := `INSERT INTO classifications (email_content, classification_result) VALUES (?, ?)`
    _, err := db.Exec(insertQuery, emailContent, classificationResult)
    return err
}

func GetClassifications() ([]map[string]interface{}, error) {
    rows, err := db.Query("SELECT id, email_content, classification_result, created_at FROM classifications")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var classifications []map[string]interface{}
    for rows.Next() {
        var id int
        var emailContent, classificationResult, createdAt string
        if err := rows.Scan(&id, &emailContent, &classificationResult, &createdAt); err != nil {
            return nil, err
        }
        classifications = append(classifications, map[string]interface{}{
            "id":                   id,
            "email_content":        emailContent,
            "classification_result": classificationResult,
            "created_at":          createdAt,
        })
    }
    return classifications, nil
}

func CloseDatabase() error {
    return db.Close()
}