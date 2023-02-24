package minimumOperations

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)

	// 本质上就是统计非0的 不同的数字个数
	flag := 0
	ans := 0
	for _, v := range nums {
		if flag == v {
			continue
		}
		flag = v
		ans++
	}
	return ans
}
