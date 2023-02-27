package movesToMakeZigzag

func movesToMakeZigzag(nums []int) int {
	// 要么所有的基数下标 都小于 左右的值
	// 要么所有的偶数下标 都小于 左右的值
	var help = func(pos int) (res int) {
		for ; pos < len(nums); pos += 2 {
			a := 0
			if pos-1 >= 0 {
				a = max(a, nums[pos]-nums[pos-1]+1)
			}

			if pos+1 < len(nums) {
				a = max(a, nums[pos]-nums[pos+1]+1)
			}
			res += a
		}
		return res
	}

	return min(help(0), help(1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
