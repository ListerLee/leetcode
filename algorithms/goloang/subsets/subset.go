package subsets

func subsets(nums []int) [][]int {

	ans := [][]int{}
	a := []int{}
	var dfs func(cur int)
	dfs = func(cur int) {

		if cur == len(nums) {
			ans = append(ans, append([]int(nil), a...))
		}
		a = append(a, nums[cur])
		dfs(cur + 1)
		a = a[:len(a)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}
