package main

import (
	"FileManager/config"
	"FileManager/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()

	// Create a new Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server on port 8081
	fmt.Println("🚀 Server is running on http://localhost:8081")
	router.Run(":8081")
}
