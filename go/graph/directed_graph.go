package graph

import (
	"container/list"
	"fmt"
)

type DirectedGraph struct {
	adjList [][]int
}

func (g *DirectedGraph) AddEdge(source, target int) {

	g.adjList[source] = append(g.adjList[source], target)
}

func (g *DirectedGraph) DFS() {

	stack := list.New()
	visited := make([]bool, len(g.adjList))

	for vertex := range g.adjList {
		if !visited[vertex] {
			stack.PushBack(vertex)
			visited[vertex] = true
			for stack.Len() > 0 {
				last := stack.Remove(stack.Back()).(int)
				fmt.Printf("%d ", last)

				for _, edge := range g.adjList[last] {
					if !visited[edge] {
						stack.PushBack(edge)
						visited[edge] = true
					}
				}
			}
		}
	}
}

func (g *DirectedGraph) TopologicalSort() {

	stack := list.New()
	answer := list.New()
	visited := make([]bool, len(g.adjList))

	for vertex := range g.adjList {
		if !visited[vertex] {
			stack.PushBack(vertex)
			visited[vertex] = true
			for stack.Len() > 0 {
				last := stack.Remove(stack.Back()).(int)

				for _, edge := range g.adjList[last] {
					if !visited[edge] {
						stack.PushBack(edge)
						visited[edge] = true
					}
				}
				answer.PushBack(last)
			}
		}
	}

	for answer.Len() > 0 {
		fmt.Printf("%d ", answer.Remove(answer.Back()).(int))
	}
}

// Topological sort Kahn's Algorithm
func (g *DirectedGraph) TopologicalSortKahn() {

	inDegree := make([]int, len(g.adjList))
	for _, edges := range g.adjList {
		for _, edge := range edges {
			inDegree[edge]++
		}
	}

	queue := list.New()
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue.PushBack(vertex)
		}
	}

	for queue.Len() > 0 {
		vertex := queue.Remove(queue.Front()).(int)
		fmt.Printf("%d ", vertex)
		for _, edge := range g.adjList[vertex] {
			inDegree[edge]--
			if inDegree[edge] == 0 {
				queue.PushBack(edge)
			}
		}
	}
}

// Time Complexity: O(V+E)
// Space Complexity: O(V)
