package threeSum

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)

	ans = [][]int{}

	// 遍历第一个数
	for i := 0; i < len(nums); i++ {
		// 如果数字有重复 直接跳过
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		num1 := nums[i]

		// 遍历第二个数
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[j-1] && j-1 != i {
				continue
			}
			num2 := nums[j]

			// 遍历第三个数
			for k := j + 1; k < len(nums); k++ {
				num3 := nums[k]

				// 三个数相加为0 则符合追加到切片中
				if num1+num2+num3 == 0 {
					ans = append(ans, []int{num1, num2, num3})
					break
				}

				// num3 为后续元素中最小值, 已经大于0 就不必比较后面的元素
				if num1+num2+num3 > 0 {
					break
				}
			}
		}
	}
	return
}
