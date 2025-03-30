package graph

import "testing"

func TestDFS(suite *testing.T) {

	suite.Parallel()
	graph := &Graph{
		adjList: make([][]int, 10),
	}

	graph.AddEdge(1, 5)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(5, 8)
	graph.AddEdge(2, 6)
	graph.AddEdge(8, 9)
	graph.AddEdge(9, 2)

	suite.Run("print dfs", func(t *testing.T) {
		graph.DFS(1)
	})
}
