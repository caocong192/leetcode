package mergeSimilarItems

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var m = make(map[int]int)
	for _, item := range items1 {
		m[item[0]] += item[1]
	}

	for _, item := range items2 {
		m[item[0]] += item[1]
	}

	var ans [][]int
	for key, value := range m {
		ans = append(ans, []int{key, value})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return ans
}
