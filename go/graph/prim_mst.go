package graph

// Prim's MST algorithm
// Input [u, v, w]
func (g *WeightedGraph) PrimMST() int {

	const inf = 1 << 60

	// Initialize the distance array
	dist := make([]int, len(g.adjacencyList))
	for i := range dist {
		dist[i] = inf
	}

	// Initialize the parent array
	parent := make([]int, len(g.adjacencyList))
	for i := range parent {
		parent[i] = -1
	}

	// Initialize the visited array
	visited := make([]bool, len(g.adjacencyList))

	// Initialize the min heap
	h := NewMinHeap()

	// Start from the first vertex
	dist[0] = 0
	h.Push(Pair{0, 0})

	// Loop until the heap is empty
	for h.Len() > 0 {

		// Get the top element from the heap
		top := h.Top()
		u := top.first
		h.Pop()

		// Mark the vertex as visited
		visited[u] = true

		// Loop through the adjacent vertices
		for _, edge := range g.adjacencyList[u] {
			v := edge.first
			w := edge.second

			// If the vertex is not visited and the weight is less than the current distance
			if !visited[v] && w < dist[v] {
				dist[v] = w
				parent[v] = u
				h.Push(Pair{v, w})
			}
		}
	}

	// Calculate the total weight of the MST
	total := 0
	for i := range dist {
		total += dist[i]
	}

	return total
}

// Time Complexity: O(V^2)
// Space Complexity: O(V)
