package repository

import (
	"database/sql"
	"portfolio-website/internal/models"
	"time"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) GetAllBlogs() ([]models.Blog, error) {
	rows, err := r.db.Query("SELECT id, title, date, summary, content FROM blogs ORDER BY date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var b models.Blog
		var date time.Time
		if err := rows.Scan(&b.ID, &b.Title, &date, &b.Summary, &b.Content); err != nil {
			return nil, err
		}
		// Format the date as string in YYYY-MM-DD format
		b.Date = date.Format("2006-01-02")
		blogs = append(blogs, b)
	}

	return blogs, nil
}

func (r *BlogRepository) GetBlogByID(id int) (*models.Blog, error) {
	var b models.Blog
	var date time.Time
	err := r.db.QueryRow("SELECT id, title, date, summary, content FROM blogs WHERE id = ?", id).
		Scan(&b.ID, &b.Title, &date, &b.Summary, &b.Content)
	if err != nil {
		return nil, err
	}

	// Format the date as string in YYYY-MM-DD format
	b.Date = date.Format("2006-01-02")

	return &b, nil
}

func (r *BlogRepository) CreateBlog(blog *models.Blog) (int64, error) {
	// Parse the date string from YYYY-MM-DD format
	date, err := time.Parse("2006-01-02", blog.Date)
	if err != nil {
		return 0, err
	}

	result, err := r.db.Exec(
		"INSERT INTO blogs (title, date, summary, content) VALUES (?, ?, ?, ?)",
		blog.Title, date, blog.Summary, blog.Content,
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

func (r *BlogRepository) UpdateBlog(blog *models.Blog) error {
	// Parse the date string from YYYY-MM-DD format
	date, err := time.Parse("2006-01-02", blog.Date)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		"UPDATE blogs SET title = ?, date = ?, summary = ?, content = ? WHERE id = ?",
		blog.Title, date, blog.Summary, blog.Content, blog.ID,
	)
	return err
}

func (r *BlogRepository) DeleteBlog(id int) error {
	_, err := r.db.Exec("DELETE FROM blogs WHERE id = ?", id)
	return err
}
