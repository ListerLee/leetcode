package main

import (
	"fmt"
)

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

func longestPalindrome1(s string) string {
	ans := ""
	start, end := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		left1, right1 := expendAroundCenter(i, i, s)
		left2, right2 := expendAroundCenter(i, i+1, s)

		if right1-left1 > end-start {
			start = left1
			end = right1
		}

		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}

	ans = s[start : end+1]

	return ans
}

func expendAroundCenter(left, right int, s string) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
func main() {
	str := "babad"

	ans := longestPalindrome1(str)

	fmt.Println(ans)
}
