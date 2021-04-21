package main

import "fmt"

func search1(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}

	fmt.Println(search(nums, 1))
}
