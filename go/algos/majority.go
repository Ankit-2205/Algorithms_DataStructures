package algos

// MajorityElement finds the majority element in an array
// It uses the Boyer-Moore majority vote algorithm
// Time complexity: O(n)
func MajorityElement(arr []int) int {
	majority := arr[0]
	count := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == majority {
			count++
		} else {
			count--
		}

		if count == 0 {
			majority = arr[i]
			count = 1
		}
	}

	return majority
}
