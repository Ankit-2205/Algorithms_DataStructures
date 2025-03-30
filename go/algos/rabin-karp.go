package algos

func RabinKarp(text, pattern string) int {
	n, m := len(text), len(pattern)
	if n < m {
		return -1
	}

	const base = 26
	const mod = 1e9 + 7
	patternHash, textHash := 0, 0
	power := 1
	for i := 0; i < m; i++ {
		patternHash = (patternHash*base + int(pattern[i]-'a')) % mod
		textHash = (textHash*base + int(text[i]-'a')) % mod
		power = (power * base) % mod
	}

	if patternHash == textHash {
		return 0
	}

	for i := m; i < n; i++ {
		textHash = (textHash*base + int(text[i]-'a')) % mod
		textHash = (textHash - int(text[i-m]-'a')*power%mod + mod) % mod
		if textHash == patternHash {
			return i - m + 1
		}
	}

	return -1
}
