// Krukal's MST algorithm to find minimum spanning tree

package graph

import (
	"sort"
)

type DisjointSet struct {
	parents map[int]int
	sizes   map[int]int
}

func (djs *DisjointSet) Add(value int) {

	if _, ok := djs.parents[value]; !ok {
		djs.parents[value] = value
		djs.sizes[value] = 1
	}
}

func (djs *DisjointSet) FindParent(value int) int {

	if djs.parents[value] == value {
		return value
	}

	djs.parents[value] = djs.FindParent(djs.parents[value])
	return djs.parents[value]
}

func (djs *DisjointSet) Join(left, right int) {

	leftParent := djs.FindParent(left)
	rightParent := djs.FindParent(right)

	if leftParent == rightParent {
		return
	}

	if djs.sizes[leftParent] > djs.sizes[rightParent] {
		djs.parents[rightParent] = leftParent
		djs.sizes[leftParent] += djs.sizes[rightParent]
	} else {
		djs.parents[leftParent] = rightParent
		djs.sizes[rightParent] += djs.sizes[leftParent]
	}
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		parents: map[int]int{},
		sizes:   map[int]int{},
	}
}

type Pair struct {
	first  int
	second int
}

type WeightedGraph struct {
	adjacencyList [][]Pair
}

// edge = [u, v, w]
// u, v are vertices and w is weight
func (g *WeightedGraph) KruskalMST() int {

	edges := [][]int{}
	for u, pairs := range g.adjacencyList {
		for _, pair := range pairs {
			v, w := pair.first, pair.second
			edges = append(edges, []int{u, v, w})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	minimumCost := 0
	disjointSet := NewDisjointSet()
	for u := range g.adjacencyList {
		disjointSet.Add(u)
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if disjointSet.FindParent(u) != disjointSet.FindParent(v) {
			minimumCost += w
			disjointSet.Join(u, v)
		}
	}

	return minimumCost
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{}
}

// Time Complexity: O(E log E)
// Space Complexity: O(V + E)
