package algos

func LengthOfLIS(nums []int) int {
	answer := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > answer[len(answer)-1] {
			answer = append(answer, nums[i])
		} else {
			l, r := 0, len(answer)-1
			for l <= r {
				mid := (l + r) / 2
				if answer[mid] > nums[i] && (mid-1 < 0 || answer[mid-1] < nums[i]) {
					answer[mid] = nums[i]
					break
				}
				if answer[mid] > nums[i] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		}
	}

	return len(answer)
}

func LIS(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// To store the indices of the LIS
	dp := make([]int, len(nums))
	// To reconstruct the LIS
	prev := make([]int, len(nums))
	for i := range prev {
		prev[i] = -1
	}

	length := 0
	var lisEnd int

	for i := 0; i < len(nums); i++ {
		// Binary search for the position of nums[i] in dp
		low, high := 0, length
		for low < high {
			mid := (low + high) / 2
			if nums[dp[mid]] < nums[i] {
				low = mid + 1
			} else {
				high = mid
			}
		}

		// Update dp and prev arrays
		if low < len(dp) {
			dp[low] = i
		}
		if low > 0 {
			prev[i] = dp[low-1]
		}

		// If we extended the LIS, update length and lisEnd
		if low == length {
			length++
			lisEnd = i
		}
	}

	// Reconstruct the LIS subsequence
	lis := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		lis[i] = nums[lisEnd]
		lisEnd = prev[lisEnd]
	}

	return lis
}

// Time Complexity: O(n * log(n))
// Space Complexity: O(n)
