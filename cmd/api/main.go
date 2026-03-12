package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router)

	// Get port (default 8080)
	port := os.Getenv("PORT")
	fmt.Println("Port from env:", port)
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
