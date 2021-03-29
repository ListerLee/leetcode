package main

import "fmt"

func isMatch(s string, p string) bool {

	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}

		return s[i-1] == p[j-1]
	}

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(i, j) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	s := "aa"
	p := "a*"

	fmt.Println(isMatch(s, p))
}
