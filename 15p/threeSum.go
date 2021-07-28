package threeSum

import "sort"

func threeSum(nums []int) [][]int {

	// 从小到大排序 nums
	sort.Ints(nums)

	ans := make([][]int, 0)

	n := len(nums)

	// 第一次循环, 取第一个元素
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 目标值
		target := nums[first] * -1
		third := n - 1

		// 第二次循环, 获取第二个元素
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 第三次循环 获取第三个元素
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			// 如果相等 则表示没有匹配的
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return  ans
}
