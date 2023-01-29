package minimumSize

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	max := getMax(nums)

	return sort.Search(max, func(i int) bool {
		if i == 0 {
			return false
		}

		ops := 0
		for _, x := range nums {
			ops += (x - 1) / i
		}
		return ops <= maxOperations
	})
}

func getMax(nums []int) (max int) {
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return
}
