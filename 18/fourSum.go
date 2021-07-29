package fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {

	// 从小到大排序 nums
	sort.Ints(nums)

	ans := make([][]int, 0)

	n := len(nums)

	if n <= 3 {
		return ans
	}

	// 第一次循环, 取第一个元素
	for first := 0; first < n; first++ {
		// 剪枝
		if nums[first] > target && nums[first] >= 0 {
			break
		}

		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		// 第二次循环, 获取第二个元素
		for second := first + 1; second < n; second++ {
			// 剪枝
			if nums[first]+nums[second] > target && nums[second] >= 0 {
				break
			}

			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 目标值
			targetRes := target - nums[first] - nums[second]
			forth := n - 1

			// 第三次循环, 获取第三个元素
			for third := second + 1; third < n; third++ {
				// 剪枝
				if nums[first]+nums[second]+nums[third] > target && nums[third] >= 0 {
					break
				}

				// 需要和上一次枚举的数不相同
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}

				// 第四次循环 获取第四个元素
				for third < forth && nums[third]+nums[forth] > targetRes {
					forth--
				}

				// 如果相等 则表示没有匹配的
				if third == forth {
					break
				}

				if nums[third]+nums[forth] == targetRes {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[forth]})
				}
			}
		}
	}
	return ans
}
