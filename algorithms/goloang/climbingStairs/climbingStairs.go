package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	ans, dp1, dp2 := 0, 1, 1
	for i := 2; i <= n; i++ {
		ans = dp1 + dp2
		dp2 = dp1
		dp1 = ans

	}

	return ans
}

func main() {
	fmt.Println(climbStairs(3))
}
