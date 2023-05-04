package maxTotalFruits

import (
	"sort"
)

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })
	right, s := left, 0

	// 统计 left 到 startPos 的水果总数
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1]
	}
	ans := s

	// 向右 一直遍历到 startPos+k
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		// 枚举最右边的位置
		s += fruits[right][1]
		// 先向左走,再向右走. 步数不大于k 并且 先向右走,再向左走, 步数不大于K
		for startPos-fruits[left][0]*2+fruits[right][0] > k && fruits[right][0]*2-startPos-fruits[left][0] > k {
			s -= fruits[left][1]
			left++
		}
		ans = max(ans, s)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
