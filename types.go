package main

// Document comment
type Document struct {
	Language string `json:"language"`
	ID       string `json:"id"`
	Text     string `json:"text"`
}

// SentimentResponse comment
type SentimentResponse struct {
	Documents []struct {
		ID    string  `json:"id"`
		Score float64 `json:"score"`
	} `json:"documents"`
	Errors []interface{} `json:"errors"`
}
