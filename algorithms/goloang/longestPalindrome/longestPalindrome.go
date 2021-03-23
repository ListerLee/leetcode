package main

import "fmt"

func longestPalindrome(s string) string {
	ans := ""
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[j] == s[i] {
					dp[i][j] = 1
				}
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] > 0 && (l+1) > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}

func main() {
	str := "babad"

	ans := longestPalindrome(str)

	fmt.Println(ans)
}
