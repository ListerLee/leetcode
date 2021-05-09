package nextPermutation

func nextPermutation(nums []int) {
	n := len(nums)
	l := n - 2
	for l >= 0 && nums[l+1] <= nums[l] {
		l--
	}
	r := n - 1
	if l >= 0 {
		for r >= l && nums[r] <= nums[l] {
			r--
		}
		nums[r], nums[l] = nums[l], nums[r]
	}
	reverse(nums[l+1:])
	return
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
