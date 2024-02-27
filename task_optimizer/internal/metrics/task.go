package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type TaskServiceMetrics struct {
	ProcessingTime    prometheus.Summary
	BronKerboschTime  prometheus.Summary
	InputTaskListSize prometheus.Histogram
	TaskListSize      prometheus.Gauge
}

func NewTaskServiceMetrics() *TaskServiceMetrics {
	metrics := &TaskServiceMetrics{
		ProcessingTime: prometheus.NewSummary(prometheus.SummaryOpts{
			Name: "task_optimizer_processing_duration_seconds",
			Help: "Time it takes to optimize for profit the list of tasks to execute in seconds",
		}),
		BronKerboschTime: prometheus.NewSummary(prometheus.SummaryOpts{
			Name: "task_optimizer_bron_kerbosch_duration_seconds",
			Help: "Time it takes to run the BronKerbosch algorithm in the task compatibility graph",
		}),
		InputTaskListSize: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "task_optimizer_input_task_list_size",
			Help: "Input size of the task list to optimize",
		}),
		TaskListSize: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "task_optimizer_task_list_size",
			Help: "Task list size",
		}),
	}

	prometheus.MustRegister(
		metrics.ProcessingTime,
		metrics.BronKerboschTime,
		metrics.InputTaskListSize,
		metrics.TaskListSize,
	)

	return metrics
}
