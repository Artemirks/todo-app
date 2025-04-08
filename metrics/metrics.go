package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalTasks = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "todo_total_tasks",
		Help: "Общее количество задач",
	})

	CompletedTasks = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "todo_completed_tasks",
		Help: "Количество выполненных задач",
	})

	IncompleteTasks = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "todo_incomplete_tasks",
		Help: "Количество невыполненных задач",
	})

	CompletionRatio = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "todo_completion_ratio",
		Help: "Процент выполненных задач (от 0 до 1)",
	})
)

func Init() {
	prometheus.MustRegister(
		TotalTasks,
		CompletedTasks,
		IncompleteTasks,
		CompletionRatio,
	)
}
