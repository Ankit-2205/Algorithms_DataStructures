package graph

func (g *Graph) FloydWarshall() [][]int {
	dist := make([][]int, len(g.adjList))
	for i := range dist {
		dist[i] = make([]int, len(g.adjList))
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}

	for i := range g.adjList {
		dist[i][i] = 0
		for _, edge := range g.adjList[i] {
			dist[i][edge] = 1
		}
	}

	for k := range g.adjList {
		for i := range g.adjList {
			for j := range g.adjList {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
}

// Time Complexity: O(V^3)
// Space Complexity: O(V^2)
