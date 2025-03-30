package algos

import "math"

func MinWindowString(s string, t string) string {
	tmap := map[byte]int{}
	for idx := range t {
		if _, ok := tmap[t[idx]]; !ok {
			tmap[t[idx]] = 1
		} else {
			tmap[t[idx]] += 1
		}
	}

	currTracked := map[byte]int{}
	minLength := math.MaxInt
	minI, minJ := 0, 0
	i, j := 0, 0
	for i < len(s) && j <= len(s) {
		if j < i {
			j++
			continue
		}

		progressNeeded := false
		for char, count := range tmap {
			if currCount, ok := currTracked[char]; !ok || currCount < count {
				progressNeeded = true
			}
		}

		if !progressNeeded && j-i < minLength {
			minLength = j - i
			minI = i
			minJ = j
		}

		if !progressNeeded {
			currTracked[s[i]] -= 1
			i++
		} else if j < len(s) {
			currTracked[s[j]] += 1
			j++
		} else {
			j++
		}
	}

	return s[minI:minJ]
}
