package main

import "fmt"

func firstMissingPositive(nums []int) int {
	m := map[int]bool{}
	ans := 1
	n := len(nums)
	for _, v := range nums {
		m[v] = true
	}
	for i := 1; i <= n+1; i++ {
		if m[i] != true {
			ans = i
			break
		}
	}
	return ans

}

func main() {
	nums := []int{1}
	fmt.Println(firstMissingPositive(nums))
}
