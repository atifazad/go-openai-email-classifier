package models

type Email struct {
	ID             int    `json:"id"`
	Subject        string `json:"subject"`
	Content        string `json:"content"`
	Classification string `json:"classification"`
	CreatedAt      string `json:"created_at"`
}
