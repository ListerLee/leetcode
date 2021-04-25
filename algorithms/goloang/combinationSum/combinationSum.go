package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	comb := []int{}
	n := len(candidates)
	var dps func(int, int)
	dps = func(target int, index int) {
		if index == n {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, comb...))
			return
		}
		dps(target, index+1)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dps(target-candidates[index], index)
			comb = comb[:len(comb)-1]
		}
	}
	dps(target, 0)

	return ans
}

func main() {
	c := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(c, 7))
}
