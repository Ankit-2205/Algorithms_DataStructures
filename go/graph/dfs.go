package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adjList [][]int
}

func (g *Graph) AddEdge(vertex1 int, vertex2 int) {

	g.adjList[vertex1] = append(g.adjList[vertex1], vertex2)
	g.adjList[vertex2] = append(g.adjList[vertex2], vertex1)
}

func (g *Graph) DFS(start int) {

	stack := list.New()
	visited := make([]bool, len(g.adjList))
	stack.PushBack(start)
	visited[start] = true
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(int)
		fmt.Printf("%d ", node)
		for _, edge := range g.adjList[node] {
			if !visited[edge] {
				stack.PushBack(edge)
				visited[edge] = true
			}
		}
	}
}

// Time Complexity: O(V+E)
// Space Complexity: O(V)
