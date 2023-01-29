package minElements

import "math"

func minElements(nums []int, limit int, goal int) int {
	// 求和
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	// 求差值
	diff := int(math.Abs(float64(sum - goal)))

	return (diff + limit - 1) / limit
}
