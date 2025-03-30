package datastructure

type DisjointSet[T comparable] struct {
	parents map[T]T
	sizes   map[T]int
}

func (djs *DisjointSet[T]) Add(value T) {

	if _, ok := djs.parents[value]; !ok {
		djs.parents[value] = value
		djs.sizes[value] = 1
	}
}

func (djs *DisjointSet[T]) FindParent(value T) T {

	if djs.parents[value] == value {
		return value
	}

	djs.parents[value] = djs.FindParent(djs.parents[value])
	return djs.parents[value]
}

func (djs *DisjointSet[T]) Join(left, right T) {

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

func NewDisjointSet[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		parents: map[T]T{},
		sizes:   map[T]int{},
	}
}

// Time Complexity: O(1)
// Space Complexity: O(1)
