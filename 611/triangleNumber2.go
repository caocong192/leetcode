package triangleNumber

import "sort"

func triangleNumber2(nums []int) (ans int) {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 边为0的过滤
		if nums[i] == 0 {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			towSide := nums[i] + nums[j]

			for k := j + 1; k < len(nums) && towSide > nums[k]; k++ {
				ans++
			}
		}
	}
	return
}
