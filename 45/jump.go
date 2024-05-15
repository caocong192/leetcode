package jump

func jump(nums []int) (ans int) {
	end, maxPos := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			ans++
			end = maxPos
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
