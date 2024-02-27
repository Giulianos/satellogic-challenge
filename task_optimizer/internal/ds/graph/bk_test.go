package graph

import (
	"reflect"
	"task_optimizer/internal/ds/set"
	"testing"
)

type GraphImpl struct {
	weights   []float64
	adjacency [][]bool
}

func (g GraphImpl) GetNodes() set.Set[int] {
	nodes := set.Set[int]{}
	for i := range g.weights {
		nodes.Add(i)
	}
	return nodes
}

func (g GraphImpl) GetNeighbors(node int) set.Set[int] {
	neighbors := set.Set[int]{}
	for neighborIdx, isNeighbor := range g.adjacency[node] {
		if isNeighbor {
			neighbors.Add(neighborIdx)
		}
	}
	return neighbors
}

func (g GraphImpl) GetWeight(node int) float64 {
	return g.weights[node]
}

func TestBronKerbosch(t *testing.T) {
	tests := []struct {
		name       string
		graph      Graph
		wantNodes  set.Set[int]
		wantWeight float64
	}{
		{
			name: "K₀",
			graph: GraphImpl{
				weights:   []float64{},
				adjacency: [][]bool{},
			},
			wantNodes:  set.Empty[int](),
			wantWeight: 0,
		},
		{
			name: "K₄",
			graph: GraphImpl{
				weights: []float64{1, 2, 3, 4},
				adjacency: [][]bool{
					{false, true, true, true},
					{true, false, true, true},
					{true, true, false, true},
					{true, true, true, false},
				},
			},
			wantNodes:  set.Of(0, 1, 2, 3),
			wantWeight: 10,
		},
		{
			name: "C₄",
			graph: GraphImpl{
				weights: []float64{1, 2, 3, 4},
				adjacency: [][]bool{
					{false, true, false, true},
					{true, false, true, false},
					{false, true, false, true},
					{true, false, true, false},
				},
			},
			wantNodes:  set.Of[int](2, 3),
			wantWeight: 7,
		},
		{
			name: "Forest: one node with more weight than other component",
			graph: GraphImpl{
				weights: []float64{10, 2, 3, 4},
				adjacency: [][]bool{
					{false, false, false, false},
					{false, false, true, true},
					{false, true, false, true},
					{false, true, true, false},
				},
			},
			wantNodes:  set.Of[int](0),
			wantWeight: 10,
		},
		{
			name: "Forest: component of 3 nodes with more weight than the other node",
			graph: GraphImpl{
				weights: []float64{8, 2, 3, 4},
				adjacency: [][]bool{
					{false, false, false, false},
					{false, false, true, true},
					{false, true, false, true},
					{false, true, true, false},
				},
			},
			wantNodes:  set.Of[int](1, 2, 3),
			wantWeight: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliqueNodes, weight := BronKerbosch(
				set.Empty[int](),
				tt.graph.GetNodes(),
				set.Empty[int](),
				tt.graph,
			)
			if !reflect.DeepEqual(cliqueNodes, tt.wantNodes) {
				t.Errorf("BronKerbosch() got nodes = %v, want %v", cliqueNodes, tt.wantNodes)
			}
			if weight != tt.wantWeight {
				t.Errorf("BronKerbosch() got weight = %v, want %v", weight, tt.wantWeight)
			}
		})
	}
}
