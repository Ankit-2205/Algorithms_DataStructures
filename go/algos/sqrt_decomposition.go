package algos

import "math"

func SqrtDecomposition(arr []int, queries [][]int) []int {
	n := len(arr)
	blockSize := int(math.Sqrt(float64(n)))
	blocks := make([]int, blockSize)
	for i := 0; i < n; i++ {
		blocks[i/blockSize] += arr[i]
	}

	answers := make([]int, len(queries))
	for _, query := range queries {
		answers = append(answers, answerQuery(arr, blocks, query))
	}

	return answers
}

func answerQuery(arr []int, blocks []int, query []int) int {
	l, r := query[0], query[1]
	blockSize := len(blocks)
	sum := 0
	for l%blockSize != 0 && l <= r {
		sum += arr[l]
		l++
	}
	for l+blockSize <= r {
		sum += blocks[l/blockSize]
		l += blockSize
	}
	for l <= r {
		sum += arr[l]
		l++
	}
	return sum
}

// Time Complexity: O(n + q * sqrt(n))
// Space Complexity: O(sqrt(n))
