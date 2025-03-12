package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/SuranaAnmol/Alle-Backend-Assignment-Go/controllers"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTaskByID)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
