package api

import (
	"portfolio-website/internal/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	projectRepo *repository.ProjectRepository,
	blogRepo *repository.BlogRepository,
	contactRepo *repository.ContactRepository,
	jwtSecret string,
) {
	handlers := NewHandlers(projectRepo, blogRepo, contactRepo, jwtSecret)

	// Public routes
	public := router.Group("/api")
	{
		public.GET("/ping", handlers.Ping)
		public.GET("/projects", handlers.GetProjects)
		public.GET("/projects/:id", handlers.GetProject)
		public.GET("/blogs", handlers.GetBlogs)
		public.GET("/blogs/:id", handlers.GetBlog)
		public.POST("/contact", handlers.CreateContact)
		public.POST("/login", handlers.Login)
	}

	// Protected routes (require authentication)
	protected := router.Group("/api")
	protected.Use(AuthMiddleware(jwtSecret))
	{
		// Project CRUD
		protected.POST("/projects", handlers.CreateProject)
		protected.PUT("/projects/:id", handlers.UpdateProject)
		protected.DELETE("/projects/:id", handlers.DeleteProject)

		// Blog CRUD
		protected.POST("/blogs", handlers.CreateBlog)
		protected.PUT("/blogs/:id", handlers.UpdateBlog)
		protected.DELETE("/blogs/:id", handlers.DeleteBlog)

		// Contact CRUD
		protected.GET("/contacts", handlers.GetContacts)
		protected.GET("/contacts/:id", handlers.GetContact)
		protected.DELETE("/contacts/:id", handlers.DeleteContact)
	}
}
