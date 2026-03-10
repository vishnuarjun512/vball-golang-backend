package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"vball/internal/database"
	"vball/internal/routes"
)

func main() {

	// Set Gin to release mode (removes debug spam)
	gin.SetMode(gin.ReleaseMode)

	// Connect to database
	database.Connect()

	// Setup router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Get port (default 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)

	// Start server
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
