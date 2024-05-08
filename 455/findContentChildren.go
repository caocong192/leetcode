package findContentChildren

import (
	"sort"
)

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)

	// 双指针
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return
}
