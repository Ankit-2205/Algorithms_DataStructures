package graph

func (g *Graph) BFSWithDistance(start int) []int {
	visited := make([]bool, len(g.adjList))
	distance := make([]int, len(g.adjList))
	visited[start] = true
	queue := []int{start}

	for len(queue) > 0 {
		last := queue[0]
		queue = queue[1:]

		for _, edge := range g.adjList[last] {
			if !visited[edge] {
				queue = append(queue, edge)
				visited[edge] = true
				distance[edge] = distance[last] + 1
			}
		}
	}
	return distance
}

// Time complexity: O(V + E)
// Space complexity: O(V)
