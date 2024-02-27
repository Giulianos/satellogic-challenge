package graph

import "task_optimizer/internal/ds/set"

type Graph interface {
	GetNodes() set.Set[int]
	GetNeighbors(node int) set.Set[int]
	GetWeight(node int) float64
}
