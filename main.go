package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"contacts-api/internal/database"
	"contacts-api/internal/handlers"
	"contacts-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func main() {
	// Initialize database with pure Go SQLite driver (modernc.org/sqlite)
	// Use database/sql directly with modernc driver, then pass to GORM
	sqlDB, err := sql.Open("sqlite", "file:contacts.db?cache=shared")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Contact{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize database layer
	dbLayer := database.NewContactDB(db)

	// Initialize handlers
	contactHandler := handlers.NewContactHandler(dbLayer)

	// Setup router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Root - simple welcome page for browsers
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, `
<!DOCTYPE html>
<html>
<head>
    <title>Contacts API</title>
    <style>
        body { font-family: system-ui, sans-serif; max-width: 800px; margin: 2rem auto; padding: 0 1rem; }
        h1 { color: #333; }
        .endpoint { background: #f5f5f5; padding: 1rem; margin: 0.5rem 0; border-radius: 4px; font-family: monospace; }
        .method { color: #2563eb; font-weight: bold; }
        .get { color: #16a34a; }
        .post { color: #ea580c; }
        .put { color: #ca8a04; }
        .delete { color: #dc2626; }
        code { background: #e5e5e5; padding: 0.2rem 0.4rem; border-radius: 3px; }
        a { color: #2563eb; }
    </style>
</head>
<body>
    <h1>📇 Contacts API</h1>
    <p>REST API for managing contacts. Built with Go, Gin, GORM, and SQLite.</p>

    <h2>Endpoints</h2>
    <div class="endpoint"><span class="method get">GET</span> <code>/health</code> — Health check</div>
    <div class="endpoint"><span class="method get">GET</span> <code>/api/v1/contacts</code> — List all contacts (add <code>?q=search</code> to filter)</div>
    <div class="endpoint"><span class="method get">GET</span> <code>/api/v1/contacts/:id</code> — Get single contact</div>
    <div class="endpoint"><span class="method post">POST</span> <code>/api/v1/contacts</code> — Create contact</div>
    <div class="endpoint"><span class="method put">PUT</span> <code>/api/v1/contacts/:id</code> — Update contact</div>
    <div class="endpoint"><span class="method delete">DELETE</span> <code>/api/v1/contacts/:id</code> — Delete contact</div>

    <h2>Example</h2>
    <pre><code>curl http://localhost:8080/api/v1/contacts</code></pre>
    <p><a href="/health">Check health</a> | <a href="/api/v1/contacts">List contacts</a></p>
</body>
</html>
`)
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		contacts := api.Group("/contacts")
		{
			contacts.GET("", contactHandler.GetContacts)
			contacts.GET("/:id", contactHandler.GetContact)
			contacts.POST("", contactHandler.CreateContact)
			contacts.PUT("/:id", contactHandler.UpdateContact)
			contacts.DELETE("/:id", contactHandler.DeleteContact)
		}
	}

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}