package graph

import "testing"

func TestTopologicalSort(suite *testing.T) {

	graph := &DirectedGraph{
		adjList: make([][]int, 10),
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 7)
	graph.AddEdge(5, 6)
	graph.AddEdge(6, 7)
	graph.AddEdge(4, 9)
	graph.AddEdge(1, 9)

	suite.Run("test topological sort order", func(t *testing.T) {
		graph.TopologicalSort()
	})
}
