package nextGreaterElements

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	// 双循环,  应该用单调栈处理-todo
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < n*2-1; j++ {
			if nums[j%len(nums)] > nums[i] {
				ans[i] = nums[j%len(nums)]
				break
			}
		}
	}
	return ans
}
