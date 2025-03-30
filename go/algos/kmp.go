package algos

func KMP(text, pattern string) int {
	n, m := len(text), len(pattern)
	if n < m {
		return -1
	}

	// Compute LPS array
	// LPS[i] = the longest proper prefix of pattern[0..i] which is also a suffix of pattern[0..i]
	// For example, if pattern = "abcab", then LPS = [0, 0, 0, 1, 2]
	// We will use this array to avoid unnecessary comparisons
	// For example, if we have matched "abc" in the text and the next character is "d",
	// we can skip comparing "a" and "d" and directly compare "b" and "d"
	// This is because we know that the longest proper prefix of pattern[0..2] which is also a suffix of pattern[0..2] is "ab"
	// So, we can skip comparing "a" and "d" and directly compare "b" and "d"
	// This is the main idea behind the Knuth-Morris-Pratt (KMP) algorithm
	lps := computeLPS(pattern)
	j := 0
	for i := 0; i < n; i++ {

		// If text[i] != pattern[j], then we need to backtrack
		// We do this by setting j = lps[j-1]
		for j > 0 && text[i] != pattern[j] {
			j = lps[j-1]
		}

		// If text[i] == pattern[j], then we increment j
		// This is because we have found a match
		if text[i] == pattern[j] {
			j++
		}

		// If j == m, then we have found the pattern in the text
		// We return the starting index of the pattern in the text
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

func computeLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	j := 0

	// Compute LPS array
	// LPS[i] = the longest proper prefix of pattern[0..i] which is also a suffix of pattern[0..i]
	// For example, if pattern = "abcab", then LPS = [0, 0, 0, 1, 2]
	for i := 1; i < m; i++ {

		// If pattern[i] != pattern[j], then we need to backtrack
		// We do this by setting j = lps[j-1]
		// This is because lps[j-1] is the length of the longest proper prefix of pattern[0..j-1] which is also a suffix of pattern[0..j-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}

		// If pattern[i] == pattern[j], then we increment j
		// This is because we have found a match
		if pattern[i] == pattern[j] {
			j++
		}

		// Set lps[i] = j
		// This is because lps[i] is the length of the longest proper prefix of pattern[0..i] which is also a suffix of pattern[0..i]
		lps[i] = j
	}

	return lps
}
