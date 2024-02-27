package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"task_optimizer/internal/controller"
	"task_optimizer/internal/handler"
	"task_optimizer/internal/metrics"
	"task_optimizer/internal/service"
	"time"
)

func main() {
	logFile, err := os.OpenFile("/logs/task_optimizer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("error opening log file")
	}
	defer logFile.Close()
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(logFile).With().Timestamp().Logger()

	taskService := service.NewTaskService(metrics.NewTaskServiceMetrics())
	taskController := controller.NewTaskController(taskService)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("GET /tasks", handler.ToLoggedHandlerFunc(taskController.ListTasks))
	http.HandleFunc("POST /tasks", handler.ToLoggedHandlerFunc(taskController.AddTasks))
	http.HandleFunc("POST /tasks/execution", handler.ToLoggedHandlerFunc(taskController.GetHigherProfitTasks))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Err(err).Send()
	}
}
