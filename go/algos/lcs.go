package algos

func LongestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text1)+1)
	for i := 1; i <= len(text2); i++ {
		lastVal := 0
		for idx := 1; idx <= len(text1); idx++ {
			temp := dp[idx]
			dp[idx] = max(dp[idx-1], dp[idx])
			if text1[idx-1] == text2[i-1] {
				dp[idx] = max(dp[idx], lastVal+1)
			}
			lastVal = temp
		}
	}

	return dp[len(text1)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
