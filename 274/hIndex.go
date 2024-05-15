package hIndex

import "sort"

func hIndex(citations []int) (ans int) {
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > ans; i-- {
		ans++
	}
	return
}
