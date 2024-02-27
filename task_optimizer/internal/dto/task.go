package dto

import (
	"task_optimizer/internal/ds/set"
	"task_optimizer/internal/model"
)

type Task struct {
	Name      string   `json:"name"`
	Resources []string `json:"resources"`
	Profit    float64  `json:"profit"`
}

func (t Task) ToModel() model.Task {
	return model.Task{
		Name:      t.Name,
		Resources: set.Of(t.Resources...),
		Profit:    t.Profit,
	}
}

func TaskFromModel(task model.Task) Task {
	return Task{
		Name:      task.Name,
		Resources: task.Resources.Slice(),
		Profit:    task.Profit,
	}
}
