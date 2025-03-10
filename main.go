package main

import (
	"fmt"
	"log"
	"todo-app/database"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB(false)
	fmt.Println("Успешное подключение к базе данных!")

	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}
