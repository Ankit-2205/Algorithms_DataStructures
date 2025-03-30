package graph

import (
	"container/heap"
)

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].second < h[j].second
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(pair any) {
	*h = append(*h, pair.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	pair := old[n-1]
	*h = old[:n-1]
	return pair
}

func (h *MinHeap) Top() Pair {
	return (*h)[0]
}

func (h *MinHeap) Update(pair Pair, index int) {
	(*h)[index] = pair
	heap.Fix(h, index)
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	heap.Init(h)
	return h
}

// input: [[u, v, w]]
func (g *WeightedGraph) AddEdge(u, v, w int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], Pair{v, w})
	g.adjacencyList[v] = append(g.adjacencyList[v], Pair{u, w})
}

// Djikstra's Algo
func (g *WeightedGraph) Djikstra(start int) []int {
	const inf = 1 << 60
	dist := make([]int, len(g.adjacencyList))
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0

	pq := NewMinHeap()
	heap.Push(pq, Pair{start, 0})

	for pq.Len() > 0 {
		pair := heap.Pop(pq).(Pair)
		u, d := pair.first, pair.second
		if d > dist[u] {
			continue
		}

		for _, edge := range g.adjacencyList[u] {
			v, w := edge.first, edge.second
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				heap.Push(pq, Pair{v, dist[v]})
			}
		}
	}

	return dist
}

// Time Complexity: O((V+E)log(V))
// Space Complexity: O(V+E)
