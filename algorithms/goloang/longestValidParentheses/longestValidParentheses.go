package main

import "fmt"

func longestValidParentheses(s string) int {
	ans := 0
	n := len(s)
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		if s[i-1] == ')' {
			if s[i-2] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				if i-1-dp[i-1] > 0 && s[i-1-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				}
			}
			ans = max(ans, dp[i])
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	s := "())"
	fmt.Println(longestValidParentheses(s))
}
