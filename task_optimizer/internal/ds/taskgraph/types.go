package taskgraph

import (
	"task_optimizer/internal/ds/set"
	"task_optimizer/internal/model"
)

type TaskCompatibilityGraph struct {
	tasks            []model.Task
	compatibilityMap map[int]set.Set[int]
}

func (t TaskCompatibilityGraph) GetNodes() set.Set[int] {
	nodes := set.Empty[int]()
	for i := range t.tasks {
		nodes.Add(i)
	}

	return nodes
}

func (t TaskCompatibilityGraph) GetNeighbors(node int) set.Set[int] {
	return t.compatibilityMap[node]
}

func (t TaskCompatibilityGraph) GetWeight(node int) float64 {
	if node >= 0 && node < len(t.tasks) {
		return t.tasks[node].Profit
	}
	return 0
}

func (t TaskCompatibilityGraph) GetTasksFromNodes(nodes set.Set[int]) []model.Task {
	tasks := make([]model.Task, 0, len(nodes))
	for node := range nodes {
		if node >= 0 && node < len(t.tasks) {
			tasks = append(tasks, t.tasks[node])
		}
	}

	return tasks
}

func BuildCompatibilityGraph(tasks []model.Task) TaskCompatibilityGraph {
	cGraph := TaskCompatibilityGraph{
		tasks:            tasks[:],
		compatibilityMap: make(map[int]set.Set[int], len(tasks)),
	}
	for i := range tasks {
		cGraph.compatibilityMap[i] = set.Set[int]{}
	}
	for i, task := range tasks {
		for j := i; j < len(tasks); j++ {
			otherTask := tasks[j]
			if task.IsCompatible(otherTask) {
				cGraph.compatibilityMap[i].Add(j)
				cGraph.compatibilityMap[j].Add(i)
			}
		}
	}

	return cGraph
}
