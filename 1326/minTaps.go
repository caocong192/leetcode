package minTaps

func minTaps(n int, ranges []int) int {

	// rightMost 记录当前点 能到最远的右端
	rightMost := make([]int, n+1)
	for i, v := range ranges {
		left := max(i-v, 0)
		rightMost[left] = max(rightMost[left], i+v)
	}

	ans := 0
	currRight := 0
	nextRight := 0
	for i, v := range rightMost {
		nextRight = max(nextRight, v)

		// 如果 nextRight 已经大于 n, 直接返回结果
		if nextRight >= n {
			return ans + 1
		}

		// 如果到达 当前的最右点 currRight
		if i == currRight {
			// 并且 下一个 nextRight 没有了, 则证明 无法继续走下去
			if i == nextRight {
				return -1
			}

			// 刷新当前最右端
			currRight = nextRight
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
