package handlers

import (
	"fmt"
	"net/http"
	"todo-app/database"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Получить все задачи
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить задачи"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Получить задачу по ID
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Задача с ID %s не найдена", id)})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Создать новую задачу
func CreateTask(c *gin.Context) {

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	validate := validator.New()

	if err := validate.Struct(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации", "details": err.Error()})
		return
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать задачу"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// Обновить задачу
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Задача с ID %s не найдена", id)})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	validate := validator.New()

	if err := validate.Struct(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации", "details": err.Error()})
		return
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить задачу"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Удалить задачу
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Задача с ID %s не найдена", id)})
		return
	}
	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить задачу"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Задача успешно удалена"})
}
