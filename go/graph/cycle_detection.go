package graph

func (g *DirectedGraph) HasCycleKahns() bool {
	inDegree := make([]int, len(g.adjList))

	for _, edges := range g.adjList {
		for _, edge := range edges {
			inDegree[edge]++
		}
	}

	queue := []int{}
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}

	count := 0
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		count++

		for _, edge := range g.adjList[vertex] {
			inDegree[edge]--
			if inDegree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}

	return count != len(g.adjList)
}

func (g *Graph) HasCycle() bool {
	visited := make([]bool, len(g.adjList))
	recStack := make([]bool, len(g.adjList))

	for vertex := range g.adjList {
		if g.HasCycleUtil(vertex, visited, recStack) {
			return true
		}
	}
	return false
}

func (g *Graph) HasCycleUtil(vertex int, visited, recStack []bool) bool {
	if !visited[vertex] {
		visited[vertex] = true
		recStack[vertex] = true

		for _, edge := range g.adjList[vertex] {
			if !visited[edge] && g.HasCycleUtil(edge, visited, recStack) {
				return true
			} else if recStack[edge] {
				return true
			}
		}
	}
	recStack[vertex] = false
	return false
}

func HasCycleDSU(edges [][]int) bool {
	parent := make([]int, len(edges))
	rank := make([]int, len(edges))

	for i := range parent {
		parent[i] = i
		rank[i] = 0
	}

	for _, edge := range edges {
		x := find(parent, edge[0])
		y := find(parent, edge[1])

		if x == y {
			return true
		}
		union(parent, rank, x, y)
	}
	return false
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent []int, rank []int, x, y int) {
	xRoot := find(parent, x)
	yRoot := find(parent, y)

	if rank[xRoot] < rank[yRoot] {
		parent[xRoot] = yRoot
	} else if rank[xRoot] > rank[yRoot] {
		parent[yRoot] = xRoot
	} else {
		parent[yRoot] = xRoot
		rank[xRoot]++
	}
}

// Time Complexity: O(V+E)
// Space Complexity: O(V)
