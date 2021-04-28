package main

import "fmt"

func firstMissingPositive1(nums []int) int {
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

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		v := abs(nums[i])
		if v <= n {
			nums[v-1] = -abs(nums[v-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	nums := []int{1}
	fmt.Println(firstMissingPositive(nums))
}
