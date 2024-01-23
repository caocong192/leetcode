package alternatingSubarray

func alternatingSubarray(nums []int) int {
	ans := 0
	res := 0
	value := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]-value {
			res++
			value *= -1
		} else if res > 0 {
			ans = max(ans, res+1)
			res = 0
			value = 1
			i--
		}
	}
	if res > 0 {
		ans = max(ans, res+1)
	}
	if ans == 0 {
		return -1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
