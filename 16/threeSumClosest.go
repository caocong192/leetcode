package threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	//
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0] + nums[1] + nums[2]

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := n - 1
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 第三次循环 获取第三个元素
			for second < third && nums[first]+nums[second]+nums[third] >= target {
				ans = minDiffValue(nums[first]+nums[second]+nums[third], ans, target)
				if ans == target {
					return ans
				}
				third--
			}
			if second == third {
				break
			}
			ans = minDiffValue(nums[first]+nums[second]+nums[third], ans, target)
		}
	}
	return ans
}

func minDiffValue(a, b, target int) int {
	if math.Abs(float64(a)-float64(target)) > math.Abs(float64(b)-float64(target)) {
		return b
	} else {
		return a
	}
}
