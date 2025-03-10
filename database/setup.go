package database

import (
	"fmt"
	"log"
	"os"
	"todo-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(isTest bool) {

	var err error
	if isTest {
		err = godotenv.Load("../.env")
	} else {
		err = godotenv.Load("/app/.env")
	}

	if err != nil {
		log.Fatal("Ошибка загрузки .env файла. Убедись, что файл существует и доступен:", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных: ", err)
	}

	// Если тестируем, используем таблицу tasks_test
	if isTest {
		DB = DB.Table("tasks_test")
	}

	// Автоматически создаем таблицы на основе моделей
	DB.AutoMigrate(&models.Task{})
	fmt.Println("База данных успешно подключена!")
}
