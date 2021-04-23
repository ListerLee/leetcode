package main

import "fmt"

func searchBinary(nums []int, target int, isLarge bool) int {

	n := len(nums)
	ans := n
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target || (isLarge && nums[mid] >= target) {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	l := searchBinary(nums, target, true)
	r := searchBinary(nums, target, false) - 1

	if l <= r && r < len(nums) && nums[l] == target && nums[r] == target {
		return []int{l, r}
	} else {
		return []int{-1, -1}
	}
}

func main() {
	nums := []int{1}

	fmt.Println(searchRange(nums, 1))
}
