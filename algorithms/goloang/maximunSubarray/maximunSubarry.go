package maximunSubarray

func maxSubArray1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(dp[i], ans)
	}
	return ans
}

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		ans = max(ans, nums[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
