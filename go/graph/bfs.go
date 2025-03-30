package graph

import (
	"container/list"
)

type graph struct {
	adjList [][]int
}

func (g *graph) AddEdge(source, target int) {
	g.adjList[source] = append(g.adjList[source], target)
	g.adjList[target] = append(g.adjList[target], source)
}

func NewGraph(size int) *graph {
	return &graph{
		adjList: make([][]int, size),
	}
}

func (g *graph) BFS(start int) {
	visited := make([]bool, len(g.adjList))
	visited[start] = true
	queue := list.New()

	queue.PushBack(start)
	for queue.Len() > 0 {
		last := queue.Remove(queue.Front()).(int)

		for _, edge := range g.adjList[last] {
			if !visited[edge] {
				queue.PushBack(edge)
				visited[edge] = true
			}
		}
	}
}

// Time Complexity: O(V+E)
// Space Complexity: O(V)
