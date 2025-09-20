package main

import (
	"log"
	"portfolio-website/internal/api"
	"portfolio-website/internal/config"
	"portfolio-website/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := repository.InitDB(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	projectRepo := repository.NewProjectRepository(db)
	blogRepo := repository.NewBlogRepository(db)
	contactRepo := repository.NewContactRepository(db)

	// Setup Gin router
	router := gin.Default()

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Authorization")

	// Use CORS middleware
	router.Use(cors.New(corsConfig))

	// Setup API routes
	api.SetupRoutes(router, projectRepo, blogRepo, contactRepo, cfg.JWTSecret)

	// Start server
	log.Printf("Server starting on port %s\n", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
