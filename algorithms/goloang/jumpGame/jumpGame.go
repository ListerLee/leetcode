package jumpGame

func canJump(nums []int) bool {
	end := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if i <= end {
			end = max(end, i+nums[i])
		}
		if end >= n+1 {
			return true
		}

	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
