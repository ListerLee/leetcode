package mergeintervals

import "sort"

type byStart [][]int

func (b byStart) Len() int           { return len(b) }
func (b byStart) Less(i, j int) bool { return b[i][0] < b[j][0] }
func (b byStart) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	merge := [][]int{}
	sort.Sort(byStart(intervals))

	merge = append(merge, append([]int(nil), intervals[0]...))
	for i := 1; i < n; i++ {
		if merge[len(merge)-1][1] < intervals[i][0] {
			merge = append(merge, append([]int(nil), intervals[i]...))
		} else {
			merge[len(merge)-1][1] = max(merge[len(merge)-1][1], intervals[i][1])
		}
	}
	return merge
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
