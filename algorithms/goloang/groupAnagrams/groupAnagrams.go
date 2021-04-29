package groupAnagrams

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, s := range strs {
		cnt := [26]int{}
		for _, c := range s {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], s)
	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
