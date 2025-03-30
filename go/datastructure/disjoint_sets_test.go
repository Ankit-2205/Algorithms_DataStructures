package datastructure

import (
	"fmt"
	"testing"
)

func TestDisjointSets(suite *testing.T) {

	suite.Parallel()

	suite.Run("test add and join", func(t *testing.T) {

		t.Parallel()

		disjoinSet := NewDisjointSet[int]()
		disjoinSet.Add(1)
		disjoinSet.Add(4)
		disjoinSet.Add(6)
		disjoinSet.Add(3)
		disjoinSet.Add(7)
		disjoinSet.Add(2)
		disjoinSet.Add(5)

		printSet(disjoinSet)
		disjoinSet.Join(1, 6)
		printSet(disjoinSet)
		disjoinSet.Join(3, 6)
		printSet(disjoinSet)
		disjoinSet.Join(4, 7)
		printSet(disjoinSet)
		disjoinSet.Join(5, 2)
		printSet(disjoinSet)
		disjoinSet.Join(2, 6)
		printSet(disjoinSet)
		disjoinSet.Join(3, 5)
		printSet(disjoinSet)
		disjoinSet.Join(3, 4)
		printSet(disjoinSet)
	})
}

func printSet(disjointSet *DisjointSet[int]) {

	fmt.Println("value, parent, rank")
	for value, parent := range disjointSet.parents {
		fmt.Printf("%d, %d, %d\n", value, parent, disjointSet.sizes[value])
	}
}
