package routes

import (
	"todo-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", handlers.GetTasks)         // Получить все задачи
		tasks.GET("/:id", handlers.GetTask)       // Получить задачу по ID
		tasks.POST("/", handlers.CreateTask)      // Создать задачу
		tasks.PUT("/:id", handlers.UpdateTask)    // Обновить задачу
		tasks.DELETE("/:id", handlers.DeleteTask) // Удалить задачу
	}
}
