package findValueOfPartition

import (
	"sort"
)

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)

	ans := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] < ans {
			ans = nums[i] - nums[i-1]
		}
	}
	return ans
}
