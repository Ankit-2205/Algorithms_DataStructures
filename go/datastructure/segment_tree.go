package datastructure

// We have lazy updates and range queries in this segment tree.
// The segment tree is a binary tree that is used to store the intervals or segments.

type SegmentTree struct {
	tree []int
	lazy []int
}

// NewSegmentTree creates a new segment tree.
func NewSegmentTree(n int) *SegmentTree {
	segmentTree := &SegmentTree{
		tree: make([]int, 4*n),
		lazy: make([]int, 4*n),
	}

	return segmentTree
}

// BuildTree builds the segment tree.
func (st *SegmentTree) BuildTree(arr []int, start, end, index int) {
	if start == end {
		st.tree[index] = arr[start]
		return
	}

	mid := (start + end) / 2
	st.BuildTree(arr, start, mid, 2*index+1)
	st.BuildTree(arr, mid+1, end, 2*index+2)
	st.tree[index] = st.tree[2*index+1] + st.tree[2*index+2]
}

// UpdateRange updates the range.
func (st *SegmentTree) UpdateRange(start, end, l, r, index, value int) {
	if st.lazy[index] != 0 {
		st.tree[index] += (end - start + 1) * st.lazy[index]

		if start != end {
			st.lazy[2*index+1] += st.lazy[index]
			st.lazy[2*index+2] += st.lazy[index]
		}

		st.lazy[index] = 0
	}

	if start > r || end < l {
		return
	}

	if start >= l && end <= r {
		st.tree[index] += (end - start + 1) * value

		if start != end {
			st.lazy[2*index+1] += value
			st.lazy[2*index+2] += value
		}

		return
	}

	mid := (start + end) / 2
	st.UpdateRange(start, mid, l, r, 2*index+1, value)
	st.UpdateRange(mid+1, end, l, r, 2*index+2, value)
	st.tree[index] = st.tree[2*index+1] + st.tree[2*index+2]
}

// QueryRange queries the range.
func (st *SegmentTree) QueryRange(start, end, l, r, index int) int {
	if st.lazy[index] != 0 {
		st.tree[index] += (end - start + 1) * st.lazy[index]

		if start != end {
			st.lazy[2*index+1] += st.lazy[index]
			st.lazy[2*index+2] += st.lazy[index]
		}

		st.lazy[index] = 0
	}

	if start > r || end < l {
		return 0
	}

	if start >= l && end <= r {
		return st.tree[index]
	}

	mid := (start + end) / 2
	left := st.QueryRange(start, mid, l, r, 2*index+1)
	right := st.QueryRange(mid+1, end, l, r, 2*index+2)

	return left + right
}

// Update updates the value.
func (st *SegmentTree) Update(start, end, index, value, pos int) {
	if start == end {
		st.tree[index] += value
		return
	}

	mid := (start + end) / 2
	if pos <= mid {
		st.Update(start, mid, 2*index+1, value, pos)
	} else {
		st.Update(mid+1, end, 2*index+2, value, pos)
	}

	st.tree[index] = st.tree[2*index+1] + st.tree[2*index+2]
}

// Query queries the value.
func (st *SegmentTree) Query(start, end, l, r, index int) int {
	if l <= start && r >= end {
		return st.tree[index]
	}

	if end < l || start > r {
		return 0
	}

	mid := (start + end) / 2
	left := st.Query(start, mid, l, r, 2*index+1)
	right := st.Query(mid+1, end, l, r, 2*index+2)

	return left + right
}

// Example
// segmentTree := NewSegmentTree(5)
// arr := []int{1, 3, 5, 7, 9}
// segmentTree.BuildTree(arr, 0, 4, 0)
// segmentTree.UpdateRange(0, 4, 1, 3, 0, 10)
// segmentTree.QueryRange(0, 4, 1, 3, 0)
// segmentTree.Update(0, 4, 0, 10, 2)
// segmentTree.Query(0, 4, 1, 3, 0)

// Time complexity
// BuildTree: O(n)
// UpdateRange: O(log n)
// QueryRange: O(log n)
// Update: O(log n)
// Query: O(log n)
