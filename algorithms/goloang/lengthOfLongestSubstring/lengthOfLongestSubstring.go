package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	rk := -1
	n := len(s)
	m := map[byte]int{}
	for i := 0; i < n; i++ {

		if i != 0 {
			delete(m, s[i-1])
		}

		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		if n-i <= maxLen {
			break
		}

		maxLen = max(maxLen, rk+1-i)
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
