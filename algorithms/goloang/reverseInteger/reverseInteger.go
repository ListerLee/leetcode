package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ans := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if (ans*10 > math.MaxInt32) || (ans*10 == math.MaxInt32 && pop > 7) {
			return 0
		}
		if (ans*10 < math.MinInt32) || (ans*10 == math.MinInt32 && pop < -8) {
			return 0
		}
		ans = ans*10 + pop
	}
	return ans
}

func main() {
	value := 1534236469
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(reverse(value))
}
