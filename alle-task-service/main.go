package main

import (
	"fmt"

	"github.com/SuranaAnmol/Alle-Backend-Assignment-Go/config"
	"github.com/SuranaAnmol/Alle-Backend-Assignment-Go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Starting Task Management System...")

	// Connect to MongoDB BEFORE doing anything else
	config.ConnectDB()

	router := gin.Default()
	routes.SetupRoutes(router)

	fmt.Println("🌍 Server running on http://localhost:8080")
	router.Run(":8080") // Runs the server on port 8080
}
