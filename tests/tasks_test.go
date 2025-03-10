package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"todo-app/database"
	"todo-app/handlers"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	database.InitDB(true)
	testDB = database.DB
	os.Exit(m.Run())
}

func cleanTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM tasks_test")
}

func setupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/tasks", func(c *gin.Context) { handlers.CreateTask(c) })
	router.GET("/tasks", func(c *gin.Context) { handlers.GetTasks(c) })
	router.GET("/tasks/:id", func(c *gin.Context) { handlers.GetTask(c) })
	router.PUT("/tasks/:id", func(c *gin.Context) { handlers.UpdateTask(c) })
	router.DELETE("/tasks/:id", func(c *gin.Context) { handlers.DeleteTask(c) })

	return router
}

func TestCreateTask(t *testing.T) {
	cleanTestDB(testDB)
	router := setupRouter(testDB)

	task := models.Task{Title: "Новая задача"}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseTask models.Task
	_ = json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Equal(t, task.Title, responseTask.Title)
	cleanTestDB(testDB)
}

func TestGetTasks(t *testing.T) {
	cleanTestDB(testDB)
	router := setupRouter(testDB)

	testDB.Create(&models.Task{Title: "Тестовая задача"})

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTasks []models.Task
	_ = json.Unmarshal(w.Body.Bytes(), &responseTasks)
	assert.Len(t, responseTasks, 1)
	cleanTestDB(testDB)
}

func TestUpdateTask(t *testing.T) {
	cleanTestDB(testDB)
	router := setupRouter(testDB)

	// Создаем задачу для обновления
	task := models.Task{Title: "Старая задача"}
	testDB.Create(&task)

	// Структура для обновленной задачи (не присваиваем вручную ID)
	updatedTask := models.Task{
		Title: "Обновленная задача",
	}

	// Маршалинг обновленной задачи
	taskJSON, _ := json.Marshal(updatedTask)

	// Отправка PUT-запроса для обновления задачи
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(task.ID), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверка результата
	assert.Equal(t, http.StatusOK, w.Code)

	var responseTask models.Task
	_ = json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Equal(t, updatedTask.Title, responseTask.Title)
	cleanTestDB(testDB)
}
