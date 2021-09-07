package triangleNumber

import "sort"

func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 边为0的过滤
		if nums[i] == 0 {
			continue
		}

		k := i+2
		for j := i + 1; j < len(nums); j++ {
			towSide := nums[i] + nums[j]

			if k == len(nums){}


			for ;k < len(nums) && towSide > nums[k]; k++ {
			}
			ans += k-j
		}
	}
	return
}
