package models

type Email struct {
    ID              int    `json:"id"`
    Content         string `json:"content"`
    Classification  string `json:"classification"`
    CreatedAt       string `json:"created_at"`
}