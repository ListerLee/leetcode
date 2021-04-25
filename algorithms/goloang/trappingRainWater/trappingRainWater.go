package trappingRainWater

func trap(height []int) int {
	ans := 0
	n := len(height)
	if n == 0 {
		return 0
	}
	leftheight := make([]int, n-1)
	rightheight := make([]int, n-1)

	leftheight[0] = height[0]
	rightheight[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftheight[i] = max(leftheight[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightheight[i] = max(rightheight[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		ans += (min(leftheight[i], rightheight[i]) - height[i])
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
