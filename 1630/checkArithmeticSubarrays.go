package checkArithmeticSubarrays

import (
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) (ans []bool) {
	m := len(l)
	for i := 0; i < m; i++ {
		ans = append(ans, isAP(nums[l[i]:r[i]+1]))
	}
	return
}

// 是否是等差数列 arithmetic progression
func isAP(n []int) bool {
	if len(n) <= 2 {
		return true
	}

	// 深拷贝一个 数组num
	var num []int
	for i := 0; i < len(n); i++ {
		num = append(num, n[i])
	}

	sort.Ints(num)
	diff := num[1] - num[0]
	for i := 2; i < len(num); i++ {
		if num[i]-num[i-1] != diff {
			return false
		}
	}
	return true
}
