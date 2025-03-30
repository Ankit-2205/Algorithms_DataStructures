package algos

import (
	"math"
	"sort"
)

func MosAlgorithm(arr []int, queries [][]int) []int {
	n := len(arr)
	blockSize := int(math.Sqrt(float64(n)))
	sort.Slice(queries, func(i, j int) bool {
		a, b := queries[i], queries[j]
		if a[0]/blockSize != b[0]/blockSize {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})

	l, r := 0, -1
	sum := 0
	res := make([]int, len(queries))
	for _, query := range queries {
		l, r = answerMosQuery(arr, l, r, query[0], query[1], &sum)
		res = append(res, sum)
	}

	return res
}

func answerMosQuery(arr []int, l, r, ql, qr int, sum *int) (int, int) {
	for r < qr {
		r++
		*sum += arr[r]
	}
	for r > qr {
		*sum -= arr[r]
		r--
	}
	for l < ql {
		*sum -= arr[l]
		l++
	}
	for l > ql {
		l--
		*sum += arr[l]
	}
	return l, r
}

// Time Complexity: O((n + q) * sqrt(n) * log(sqrt(n)))
// Space Complexity: O(sqrt(n))
