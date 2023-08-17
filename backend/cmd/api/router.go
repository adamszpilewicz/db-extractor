// router.go
package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type router struct {
	gin *gin.Engine
	app *application
}

func newRouter(app *application) *router {
	r := &router{
		gin: gin.Default(),
		app: app,
	}
	return r
}

func (r *router) setupRoutes() {
	// Setup CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Adjust this to your needs.
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Link"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.gin.Use(cors.New(config))

	// Define routes
	r.gin.GET("/users/login", r.app.Login)
	r.gin.POST("/users/login", r.app.Login)
	r.gin.POST("/schemas", r.app.Schemas)
	r.gin.POST("/schemas/:tableName", r.app.TableColumns)
	r.gin.POST("/download/:tableName", r.app.DownloadTableCSV)
	r.gin.POST("/custom-query", r.app.ExecuteCustomQuery)
	r.gin.POST("/views", r.app.AppViews)
	r.gin.POST("/dbinfo", r.app.GetAppInfo)
	r.gin.POST("/db-stats", r.app.DBStatistics)

}

func (r *router) start(port string) error {
	r.app.infoLog.Printf("API listening on port %s", port)
	return r.gin.Run(port)
}

// jsonResponse is the type used for generic JSON responses
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Login is the handler used to attempt to log a user into the api
func (app *application) Login(c *gin.Context) {
	var creds Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid JSON",
		})
		return
	}

	// Create Database Repo
	repo := NewDatabaseRepo(app, creds)
	log.Printf("Trying to connect to database: %v", repo)
	err := repo.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not connect to the database",
		})
		return
	}

	// Check database connection
	err = repo.CheckConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Database connection check failed",
		})
		return
	}
	app.infoLog.Println("Database connection check passed")

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Connected succesfully to the database",
	})
}

// Schemas is the handler used to retrieve available schemas and tables
func (app *application) Schemas(c *gin.Context) {
	var creds Credentials

	// Bind the POST body to the Credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are required"})
		return
	}

	repo := NewDatabaseRepo(app, creds)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		return
	}

	schemas, err := repo.DiscoverSchemas()
	if err != nil {
		app.errorLog.Println("Could not retrieve schemas:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve schemas",
		})
		return
	}

	log.Printf("Schemas retrieved successfully: %v", schemas)

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Schemas retrieved successfully",
		"data":    schemas,
	})
}

// TableColumns is the handler used to retrieve columns of a specific table
func (app *application) TableColumns(c *gin.Context) {
	var creds Credentials

	// Bind the POST body to the Credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are required"})
		return
	}

	repo := NewDatabaseRepo(app, creds)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		return
	}

	tableName := c.Param("tableName")
	columns, err := repo.DiscoverTableColumns(tableName)
	if err != nil {
		app.errorLog.Println("Could not retrieve columns:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve columns",
		})
		return
	}

	log.Printf("Columns for table %s retrieved successfully: %v", tableName, columns)

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Columns retrieved successfully",
		"columns": columns,
	})
}

// DownloadTableCSV is the handler used to retrieve table data and send as CSV
func (app *application) DownloadTableCSV(c *gin.Context) {
	var creds Credentials

	// Bind the POST body to the Credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are required"})
		return
	}

	repo := NewDatabaseRepo(app, creds)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		return
	}

	tableName := c.Param("tableName")
	rows, err := repo.FetchTableData(tableName) // You need to implement FetchTableData in your repo.
	if err != nil {
		app.errorLog.Println("Could not retrieve rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve rows",
		})
		return
	}

	// Convert the rows to CSV.
	csvData, err := ConvertRowsToCSV(rows) // Handle both returned values
	if err != nil {
		app.errorLog.Println("Error while converting rows to CSV:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to convert rows to CSV",
		})
		return
	}

	repo.DB.Close()
	// Convert the byte slice to a string before sending.
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", tableName))
	c.String(http.StatusOK, string(csvData))
}

type CustomQueryRequest struct {
	Credentials Credentials `json:"credentials"`
	Query       string      `json:"query"`
}

func (app *application) ExecuteCustomQuery(c *gin.Context) {
	var request CustomQueryRequest

	// Bind the POST body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or missing fields"})
		return
	}

	repo := NewDatabaseRepo(app, request.Credentials)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		return
	}

	// Use the provided SQL query
	result, err := repo.ExecuteQuery(request.Query)
	if err != nil {
		app.errorLog.Println("Error executing custom query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  result,
	})
}

func (app *application) AppViews(c *gin.Context) {
	var creds Credentials

	// Bind the POST body to the Credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are required"})
		return
	}

	repo := NewDatabaseRepo(app, creds)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		return
	}

	views, err := repo.FetchAppViews()
	if err != nil {
		app.errorLog.Println("Could not retrieve app views:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve app views",
		})
		return
	}

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  views,
	})
}

func (app *application) GetAppInfo(c *gin.Context) {
	var creds Credentials

	// Bind the POST body to the Credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are required"})
		return
	}

	// Use NewDatabaseRepo and Connect using the provided creds
	repo := NewDatabaseRepo(app, creds)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not connect to the database",
		})
		return
	}

	// Fetch the App Info
	info, err := repo.GetAppInfo()
	if err != nil {
		app.errorLog.Println("Could not retrieve app info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve app info",
		})
		return
	}

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  info,
	})
}

func (app *application) DBStatistics(c *gin.Context) {
	var creds Credentials

	// Bind the POST body to the Credentials struct
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credentials are required"})
		return
	}

	// Use NewDatabaseRepo and Connect using the provided creds
	repo := NewDatabaseRepo(app, creds)
	err := repo.Connect()
	if err != nil {
		log.Printf("Could not connect to the database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not connect to the database",
		})
		return
	}

	// Fetch the Database Statistics
	stats, err := repo.FetchDBStatistics()
	if err != nil {
		app.errorLog.Println("Could not retrieve database statistics:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve database statistics",
		})
		return
	}

	repo.DB.Close()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  stats,
	})
}
