package datastructure

import "testing"

func TestPriorityQueue(suite *testing.T) {

	suite.Parallel()

	suite.Run("test insert", func(t *testing.T) {
		t.Parallel()

		pq := NewPriorityQueue()
		pq.Insert(1, "A")
		pq.Insert(10, "J")
		pq.Insert(3, "C")
		pq.Insert(6, "F")
		pq.Insert(0, "_")
		pq.Insert(8, "H")
		pq.Insert(9, "I")
		pq.Insert(-3, "-")
		pq.Insert(5, "E")

		pq.PrintQueue()

		pq.Pop()
		pq.Pop()

		pq.PrintQueue()

		pq.Insert(7, "G")
		pq.Insert(4, "D")

		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()

		pq.Pop()
		pq.PrintQueue()
	})
}

// Time Complexity: O(n)
// Space Complexity: O(n)
