package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)

	for a := 0; a < n; a++ {
		if a == 0 || nums[a] != nums[a-1] {
			c := n - 1
			for b := a + 1; b < n; b++ {
				if b == a+1 || nums[b] != nums[b-1] {

					tag := -(nums[a] + nums[b])
					for c > b && nums[c] > tag {
						c--
					}
					if c <= b {
						break
					}
					if tag == nums[c] {
						ans = append(ans, []int{nums[a], nums[b], nums[c]})
					}
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{3, 0, -2, -1, 1, 2}

	fmt.Println(threeSum(nums))
}
