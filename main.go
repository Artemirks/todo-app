package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"todo-app/database"
	"todo-app/metrics"
	"todo-app/routes"
)

func main() {
	database.InitDB(false)
	fmt.Println("Успешное подключение к базе данных!")

	metrics.Init()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	go func() {
		for {
			updateMetrics()
			time.Sleep(30 * time.Second)
		}
	}()

	// Gin-приложение
	router := gin.Default()
	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
}

func updateMetrics() {
	var total, completed, incomplete int64

	database.DB.Table("tasks").Count(&total)
	database.DB.Table("tasks").Where("completed = ?", true).Count(&completed)
	database.DB.Table("tasks").Where("completed = ?", false).Count(&incomplete)

	metrics.TotalTasks.Set(float64(total))
	metrics.CompletedTasks.Set(float64(completed))
	metrics.IncompleteTasks.Set(float64(incomplete))

	if total > 0 {
		metrics.CompletionRatio.Set(float64(completed) / float64(total))
	} else {
		metrics.CompletionRatio.Set(0)
	}
}
