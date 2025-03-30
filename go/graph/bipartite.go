package graph

func (g *Graph) IsBipartite() bool {
	colors := make([]int, len(g.adjList))
	for vertex := range g.adjList {
		if colors[vertex] == 0 && !g.IsBipartiteUtil(vertex, colors, 1) {
			return false
		}
	}
	return true
}

func (g *Graph) IsBipartiteUtil(vertex int, colors []int, color int) bool {
	if colors[vertex] != 0 {
		return colors[vertex] == color
	}

	colors[vertex] = color
	for _, edge := range g.adjList[vertex] {
		if !g.IsBipartiteUtil(edge, colors, -color) {
			return false
		}
	}
	return true
}

func (g *Graph) IsBipartiteBFS() bool {
	colors := make([]int, len(g.adjList))
	for vertex := range g.adjList {
		if colors[vertex] == 0 {
			colors[vertex] = 1
			queue := []int{vertex}
			for len(queue) > 0 {
				last := queue[0]
				queue = queue[1:]
				for _, edge := range g.adjList[last] {
					if colors[edge] == 0 {
						colors[edge] = -colors[last]
						queue = append(queue, edge)
					} else if colors[edge] == colors[last] {
						return false
					}
				}
			}
		}
	}
	return true
}

// Time complexity: O(V+E)
// Space complexity: O(V)
