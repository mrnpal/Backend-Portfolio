package models

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	DemoURL     string `json:"demoUrl"`
	GithubURL   string `json:"githubUrl"`
}
