package nextPermutation

func nextPermutation(nums []int) {
	n := len(nums)
	l := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			l = i
			break
		}
	}
	if l == 0 {
		reverse(nums, 0, n-1)
		return
	}
	r := n - 1
	for i := l; i < n; i++ {
		if nums[i] <= nums[l-1] {
			r = i - 1
			break
		}
	}
	nums[l-1], nums[r] = nums[r], nums[l-1]
	reverse(nums, l, n-1)
	return
}

func reverse(nums []int, l, r int) {
	n := r - l + 1
	for i := 0; i < n/2; i++ {
		nums[l+i], nums[r-i] = nums[r-i], nums[l+i]
	}
}
