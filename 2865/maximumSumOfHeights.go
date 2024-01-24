package maximumSumOfHeights

func maximumSumOfHeights(maxHeights []int) (ans int64) {
	// 枚举, 便利数组, 以每个点为顶点,计算和, 求最大值
	for i := 0; i < len(maxHeights); i++ {
		pre, sum := maxHeights[i], int64(maxHeights[i])

		for j := i - 1; j >= 0; j-- {
			pre = min(pre, maxHeights[j])
			sum += int64(pre)
			//fmt.Printf("pre: %v, i: %v\n", pre, j)
		}
		suf := maxHeights[i]
		for j := i + 1; j < len(maxHeights); j++ {
			suf = min(suf, maxHeights[j])
			sum += int64(suf)
			//fmt.Printf("suf: %v, i: %v\n", suf, j)
		}
		if ans < sum {
			ans = sum
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
