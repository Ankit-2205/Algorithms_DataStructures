package datastructure

func FenwickTree(n int) []int {
	return make([]int, n+1)
}

func Build(arr []int) []int {
	fenwick := make([]int, len(arr)+1)
	for i, v := range arr {
		Update(fenwick, i+1, v)
	}
	return fenwick
}

func Update(fenwick []int, i, delta int) {
	for i < len(fenwick) {
		fenwick[i] += delta
		i += i & -i
	}
}

func Query(fenwick []int, i int) int {
	sum := 0
	for i > 0 {
		sum += fenwick[i]
		i -= i & -i
	}
	return sum
}

// Time Complexity: O(log(n))
// Space Complexity: O(n)
