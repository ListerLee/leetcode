package containerWithMostWater

func maxArea(height []int) int {
	var ans int = 0
	x, y := 0, len(height)-1

	for y > x {
		c := min(height[x], height[y]) * (y - x)
		ans = max(c, ans)
		if height[x] > height[y] {
			y--
		} else {
			x++
		}
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
