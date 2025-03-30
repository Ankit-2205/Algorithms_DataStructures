package algos

func ManachersAlgorithm(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// Inserting '#' characters to handle even length palindromes
	// For example, "abba" becomes "#a#b#b#a#"
	// This way, we can handle both odd and even length palindromes
	t := make([]rune, 2*n+1)
	t[0] = '#'
	for i, j := 0, 1; i < n; i, j = i+1, j+2 {
		t[j] = rune(s[i])
		t[j+1] = '#'
	}

	// Manacher's Algorithm
	// p[i] = length of the longest palindromic substring centered at i
	// c = center of the longest palindromic substring
	// r = right boundary of the longest palindromic substring
	n = 2*n + 1
	p := make([]int, n)
	c, r := 0, 0
	for i := 0; i < n; i++ {
		mirror := 2*c - i
		if i < r {
			p[i] = min(r-i, p[mirror])
		}

		a, b := i+(1+p[i]), i-(1+p[i])
		// Attempt to expand palindrome centered at i
		// If the character at a equals the character at b,
		// increment p[i] and expand the palindrome
		for a < n && b >= 0 && t[a] == t[b] {
			p[i]++
			a++
			b--
		}

		// Update c and r
		// If the palindrome centered at i expands beyond r,
		// update c and r accordingly
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	// Find the maximum length palindrome and its center
	// The length of the longest palindromic substring is p[i]
	// The center of the longest palindromic substring is i
	maxLen, centerIndex := 0, 0
	for i := 0; i < n; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	start := (centerIndex - 1 - maxLen) / 2
	return s[start : start+maxLen]
}
