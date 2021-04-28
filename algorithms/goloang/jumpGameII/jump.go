package jumpGameII

func jump(nums []int) int {
	n := len(nums)
	end := 0
	step := 0
	maxPosition := 0
	for i := 0; i < n-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
