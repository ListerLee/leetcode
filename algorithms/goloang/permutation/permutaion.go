package permutation

func permute(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	var backtrace func(int)
	backtrace = func(first int) {
		if first == n {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for i := first; i < n; i++ {
			nums[i], nums[first] = nums[first], nums[i]
			backtrace(first + 1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}
	backtrace(0)

	return ans
}
