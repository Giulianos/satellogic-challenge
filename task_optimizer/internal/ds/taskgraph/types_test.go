package taskgraph

import (
	"reflect"
	"task_optimizer/internal/ds/set"
	"task_optimizer/internal/model"
	"testing"
)

func TestBuildCompatibilityGraph(t *testing.T) {
	tests := []struct {
		name  string
		tasks []model.Task
		want  TaskCompatibilityGraph
	}{
		{
			name:  "empty list",
			tasks: []model.Task{},
			want: TaskCompatibilityGraph{
				tasks:            []model.Task{},
				compatibilityMap: map[int]set.Set[int]{},
			},
		},
		{
			name: "list with one task",
			tasks: []model.Task{
				{"task1", set.Of[string]("resource"), 1.2},
			},
			want: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
				},
			},
		},
		{
			name: "list with two compatible tasks",
			tasks: []model.Task{
				{"task1", set.Of[string]("resource1"), 1.2},
				{"task2", set.Of[string]("resource2"), 1.2},
			},
			want: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
					{"task2", set.Of[string]("resource2"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Of[int](1),
					1: set.Of[int](0),
				},
			},
		},
		{
			name: "list with two incompatible tasks",
			tasks: []model.Task{
				{"task1", set.Of[string]("resource"), 1.2},
				{"task2", set.Of[string]("resource"), 1.2},
			},
			want: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource"), 1.2},
					{"task2", set.Of[string]("resource"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
					1: set.Empty[int](),
				},
			},
		},
		{
			name: "list with two compatible tasks and one incompatible",
			tasks: []model.Task{
				{"task1", set.Of[string]("resource1"), 1.2},
				{"task2", set.Of[string]("resource2"), 1.2},
				{"task3", set.Of[string]("resource1"), 1.2},
			},
			want: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
					{"task2", set.Of[string]("resource2"), 1.2},
					{"task3", set.Of[string]("resource1"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Of[int](1),
					1: set.Of[int](0, 2),
					2: set.Of[int](1),
				},
			},
		},
		{
			name: "list with two compatible tasks with two resources",
			tasks: []model.Task{
				{"task1", set.Of[string]("resource1", "resource2"), 1.2},
				{"task2", set.Of[string]("resourceA", "resourceB", "resource2"), 1.2},
			},
			want: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1", "resource2"), 1.2},
					{"task2", set.Of[string]("resourceA", "resourceB", "resource2"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
					1: set.Empty[int](),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCompatibilityGraph(tt.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildCompatibilityGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskCompatibilityGraph_GetNodes(t1 *testing.T) {
	tests := []struct {
		name  string
		graph TaskCompatibilityGraph
		want  set.Set[int]
	}{
		{
			name: "Empty graph",
			graph: TaskCompatibilityGraph{
				tasks:            []model.Task{},
				compatibilityMap: map[int]set.Set[int]{},
			},
			want: set.Empty[int](),
		},
		{
			name: "Graph with two nodes",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource"), 1.2},
					{"task2", set.Of[string]("resource"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
					1: set.Empty[int](),
				},
			},
			want: set.Of[int](0, 1),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.graph.GetNodes(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskCompatibilityGraph_GetNeighbors(t1 *testing.T) {
	tests := []struct {
		name  string
		graph TaskCompatibilityGraph
		node  int
		want  set.Set[int]
	}{
		{
			name: "Empty graph, get neighbors of nonexistent node",
			graph: TaskCompatibilityGraph{
				tasks:            []model.Task{},
				compatibilityMap: map[int]set.Set[int]{},
			},
			node: 0,
			want: nil,
		},
		{
			name: "Graph with one node, get neighbors of node",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task", set.Of[string]("resource"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
				},
			},
			node: 0,
			want: set.Empty[int](),
		},
		{
			name: "Graph with with two connected node, get neighbors of node",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
					{"task2", set.Of[string]("resource2"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Of[int](1),
					1: set.Of[int](0),
				},
			},
			node: 0,
			want: set.Of[int](1),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.graph.GetNeighbors(tt.node); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskCompatibilityGraph_GetWeight(t1 *testing.T) {
	tests := []struct {
		name  string
		graph TaskCompatibilityGraph
		node  int
		want  float64
	}{
		{
			name: "Empty graph, get weight of nonexistent node",
			graph: TaskCompatibilityGraph{
				tasks:            []model.Task{},
				compatibilityMap: map[int]set.Set[int]{},
			},
			node: 0,
			want: 0,
		},
		{
			name: "Graph with one node, get weight of node",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task", set.Of[string]("resource"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
				},
			},
			node: 0,
			want: 1.2,
		},
		{
			name: "Graph with with two nodes, get weight of node 0",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
					{"task2", set.Of[string]("resource2"), 2.4},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Of[int](1),
					1: set.Of[int](0),
				},
			},
			node: 0,
			want: 1.2,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.graph.GetWeight(tt.node); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskCompatibilityGraph_GetTasksFromNodes(t1 *testing.T) {
	tests := []struct {
		name       string
		graph      TaskCompatibilityGraph
		nodesToGet set.Set[int]
		want       []model.Task
	}{
		{
			name: "Empty graph, get nodes not present in graph",
			graph: TaskCompatibilityGraph{
				tasks:            []model.Task{},
				compatibilityMap: map[int]set.Set[int]{},
			},
			nodesToGet: set.Of[int](0, 1, 2),
			want:       []model.Task{},
		},
		{
			name: "Graph with one node, get some nodes not present in graph",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
				},
			},
			nodesToGet: set.Of[int](0, 1, 2),
			want: []model.Task{
				{"task1", set.Of[string]("resource1"), 1.2},
			},
		},
		{
			name: "Graph with one node, get no nodes",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
				},
			},
			nodesToGet: set.Empty[int](),
			want:       []model.Task{},
		},
		{
			name: "Graph with two nodes, get one node",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
					{"task2", set.Of[string]("resource1"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
					1: set.Empty[int](),
				},
			},
			nodesToGet: set.Of[int](0),
			want: []model.Task{
				{"task1", set.Of[string]("resource1"), 1.2},
			},
		},
		{
			name: "Graph with two nodes, get both nodes",
			graph: TaskCompatibilityGraph{
				tasks: []model.Task{
					{"task1", set.Of[string]("resource1"), 1.2},
					{"task2", set.Of[string]("resource1"), 1.2},
				},
				compatibilityMap: map[int]set.Set[int]{
					0: set.Empty[int](),
					1: set.Empty[int](),
				},
			},
			nodesToGet: set.Of[int](0, 1),
			want: []model.Task{
				{"task1", set.Of[string]("resource1"), 1.2},
				{"task2", set.Of[string]("resource1"), 1.2},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.graph.GetTasksFromNodes(tt.nodesToGet); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTasksFromNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
