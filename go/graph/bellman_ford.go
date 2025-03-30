package graph

func BellmanFord(edges [][]int, n, start int) []int {
	const inf = 1e9
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0

	// Relax edges repeatedly
	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			if dist[u] != inf && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
			}
		}
	}

	// Check for negative cycles
	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			if dist[u] != inf && dist[u]+w < dist[v] {
				dist[v] = -inf
			}
		}
	}

	return dist
}

// Time Complexity: O(V*E)
// Space Complexity: O(V)
