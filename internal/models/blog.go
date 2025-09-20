package models

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"` // Ensure this is string, not time.Time
	Summary string `json:"summary"`
	Content string `json:"content"`
}
