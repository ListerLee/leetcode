package motsa

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	arrLen := len(nums1) + len(nums2)
	if arrLen%2 == 1 {
		return float64(findKInSortedArrays(nums1, nums2, arrLen/2+1))
	} else {
		return (float64(findKInSortedArrays(nums1, nums2, arrLen/2) +
			findKInSortedArrays(nums1, nums2, arrLen/2+1))) / 2
	}
}

func findKInSortedArrays(num1, num2 []int, k int) int {
	off1, off2 := 0, 0
	len1, len2 := len(num1), len(num2)

	for {

		if off1 == len1 {
			return num2[off2+k-1]
		}
		if off2 == len2 {
			return num1[off1+k-1]
		}

		if k == 1 {
			return min(num1[off1], num2[off2])
		}

		k2 := k / 2
		idx1 := min(off1+k2, len1) - 1
		idx2 := min(off2+k2, len2) - 1

		v1, v2 := num1[idx1], num2[idx2]

		if v1 > v2 {
			k -= (idx2 - off2 + 1)
			off2 = idx2 + 1
		} else {
			k -= (idx1 - off1 + 1)
			off1 = idx1 + 1
		}

	}

}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
