package controller

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"task_optimizer/internal/dto"
	"task_optimizer/internal/model"
	"task_optimizer/internal/service"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (controller *TaskController) AddTasks(w http.ResponseWriter, r *http.Request) (int, any) {
	var tasksDto []dto.Task
	err := json.NewDecoder(r.Body).Decode(&tasksDto)
	if err != nil {
		log.Err(err).Send()
		return http.StatusBadRequest, nil
	}
	tasks := make([]model.Task, 0, len(tasksDto))
	for _, taskDto := range tasksDto {
		tasks = append(tasks, taskDto.ToModel())
	}
	controller.taskService.AddTasks(tasks)
	return http.StatusOK, nil
}

func (controller *TaskController) GetHigherProfitTasks(w http.ResponseWriter, r *http.Request) (int, any) {
	taskSubset := controller.taskService.GetHigherProfitSubset()
	tasksDto := make([]dto.Task, 0, len(taskSubset))
	for _, task := range taskSubset {
		tasksDto = append(tasksDto, dto.TaskFromModel(task))
	}
	return http.StatusOK, tasksDto
}

func (controller *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) (int, any) {
	tasks := controller.taskService.ListAllTasks()
	tasksDto := make([]dto.Task, 0, len(tasks))
	for _, task := range tasks {
		tasksDto = append(tasksDto, dto.TaskFromModel(task))
	}
	return http.StatusOK, tasksDto
}
