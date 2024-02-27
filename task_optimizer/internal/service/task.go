package service

import (
	"sync"
	"task_optimizer/internal/ds/graph"
	"task_optimizer/internal/ds/set"
	"task_optimizer/internal/ds/taskgraph"
	"task_optimizer/internal/metrics"
	"task_optimizer/internal/model"
	"time"
)

type TaskService struct {
	tasksMu sync.RWMutex
	tasks   []model.Task

	metrics *metrics.TaskServiceMetrics
}

func NewTaskService(taskServiceMetrics *metrics.TaskServiceMetrics) *TaskService {
	return &TaskService{
		metrics: taskServiceMetrics,
	}
}

func (s *TaskService) AddTasks(tasks []model.Task) {
	if s.tasks == nil {
		s.tasks = make([]model.Task, 0, len(tasks))
	}
	s.tasksMu.Lock()
	for _, task := range tasks {
		s.tasks = append(s.tasks, task)
	}
	s.metrics.TaskListSize.Set(float64(len(s.tasks)))
	s.tasksMu.Unlock()
}

func (s *TaskService) ListAllTasks() []model.Task {
	s.tasksMu.RLock()
	tasks := s.tasks[:]
	s.tasksMu.RUnlock()
	return tasks
}

func (s *TaskService) GetHigherProfitSubset() []model.Task {
	startTime := time.Now()

	s.tasksMu.Lock()
	s.metrics.InputTaskListSize.Observe(float64(len(s.tasks)))
	compatibilityGraph := taskgraph.BuildCompatibilityGraph(s.tasks)
	bronKerboschStartTime := time.Now()
	taskNodesSubset, _ := graph.BronKerbosch(
		set.Empty[int](),
		compatibilityGraph.GetNodes(),
		set.Empty[int](),
		compatibilityGraph,
	)
	s.metrics.BronKerboschTime.Observe(time.Since(bronKerboschStartTime).Seconds())
	remainingNodes := compatibilityGraph.GetNodes().Difference(taskNodesSubset)
	s.tasks = compatibilityGraph.GetTasksFromNodes(remainingNodes)
	s.metrics.TaskListSize.Set(float64(len(s.tasks)))
	s.tasksMu.Unlock()

	s.metrics.ProcessingTime.Observe(time.Since(startTime).Seconds())
	return compatibilityGraph.GetTasksFromNodes(taskNodesSubset)
}
