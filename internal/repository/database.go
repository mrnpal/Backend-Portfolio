package repository

import (
	"database/sql"
	_ "log"
	"os"
	"path/filepath"
	_ "time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(databasePath string) (*sql.DB, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(databasePath), 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	// Create tables
	if err := createTables(db); err != nil {
		return nil, err
	}

	// Insert sample data if tables are empty
	if err := insertSampleData(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS projects (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            description TEXT NOT NULL,
            image_url TEXT NOT NULL,
            demo_url TEXT,
            github_url TEXT
        )`,
		`CREATE TABLE IF NOT EXISTS blogs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            date DATETIME NOT NULL,
            summary TEXT NOT NULL,
            content TEXT NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS contacts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
            message TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func insertSampleData(db *sql.DB) error {
	// Check if projects table is empty
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM projects").Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		// Insert sample projects
		projects := []struct {
			title       string
			description string
			image_url   string
			demo_url    string
			github_url  string
		}{
			{
				"E-Commerce Platform",
				"A full-featured e-commerce platform with cart functionality and payment integration.",
				"https://via.placeholder.com/600x400?text=E-Commerce",
				"https://example.com/demo1",
				"https://github.com/example/ecommerce",
			},
			{
				"Task Management App",
				"A collaborative task management application with real-time updates.",
				"https://via.placeholder.com/600x400?text=Task+Manager",
				"https://example.com/demo2",
				"https://github.com/example/taskmanager",
			},
			{
				"Weather Dashboard",
				"A weather dashboard with forecasts and interactive maps.",
				"https://via.placeholder.com/600x400?text=Weather",
				"https://example.com/demo3",
				"https://github.com/example/weather",
			},
		}

		for _, project := range projects {
			_, err := db.Exec(
				"INSERT INTO projects (title, description, image_url, demo_url, github_url) VALUES (?, ?, ?, ?, ?)",
				project.title, project.description, project.image_url, project.demo_url, project.github_url,
			)
			if err != nil {
				return err
			}
		}
	}

	// Check if blogs table is empty
	if err := db.QueryRow("SELECT COUNT(*) FROM blogs").Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		// Insert sample blogs
		blogs := []struct {
			title   string
			date    string
			summary string
			content string
		}{
			{
				"Getting Started with React",
				"2023-06-15",
				"Learn the basics of React and how to build your first application.",
				"React is a JavaScript library for building user interfaces. It allows you to create reusable UI components and efficiently update the UI when data changes.",
			},
			{
				"Building REST APIs with Go",
				"2023-07-22",
				"A comprehensive guide to creating RESTful APIs using the Go programming language.",
				"Go is an excellent choice for building high-performance REST APIs. In this article, we'll explore how to create a robust API using the Gin framework.",
			},
			{
				"Database Design Best Practices",
				"2023-08-30",
				"Essential tips and best practices for designing efficient and scalable databases.",
				"Good database design is crucial for application performance. This article covers normalization, indexing strategies, and query optimization techniques.",
			},
		}

		for _, blog := range blogs {
			_, err := db.Exec(
				"INSERT INTO blogs (title, date, summary, content) VALUES (?, ?, ?, ?)",
				blog.title, blog.date, blog.summary, blog.content,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
