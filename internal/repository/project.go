package repository

import (
	"database/sql"
	"portfolio-website/internal/models"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) GetAllProjects() ([]models.Project, error) {
	rows, err := r.db.Query("SELECT id, title, description, image_url, demo_url, github_url FROM projects ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.DemoURL, &p.GithubURL); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectRepository) GetProjectByID(id int) (*models.Project, error) {
	var p models.Project
	err := r.db.QueryRow("SELECT id, title, description, image_url, demo_url, github_url FROM projects WHERE id = ?", id).
		Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.DemoURL, &p.GithubURL)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProjectRepository) CreateProject(project *models.Project) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO projects (title, description, image_url, demo_url, github_url) VALUES (?, ?, ?, ?, ?)",
		project.Title, project.Description, project.ImageURL, project.DemoURL, project.GithubURL,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProjectRepository) UpdateProject(project *models.Project) error {
	_, err := r.db.Exec(
		"UPDATE projects SET title = ?, description = ?, image_url = ?, demo_url = ?, github_url = ? WHERE id = ?",
		project.Title, project.Description, project.ImageURL, project.DemoURL, project.GithubURL, project.ID,
	)
	return err
}

func (r *ProjectRepository) DeleteProject(id int) error {
	_, err := r.db.Exec("DELETE FROM projects WHERE id = ?", id)
	return err
}
