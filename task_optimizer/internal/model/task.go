package model

import "task_optimizer/internal/ds/set"

type Task struct {
	Name      string
	Resources set.Set[string]
	Profit    float64
}

func (task Task) IsCompatible(other Task) bool {
	for resource := range task.Resources {
		if other.Resources.Contains(resource) {
			return false
		}
	}

	return true
}
