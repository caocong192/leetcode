package stoneGame

import "math"

func stoneGame(piles []int) bool {
	var dp [500][500]int

	// k 表示 pile i 与 j 之间的间隔
	for k := 0; k < len(piles); k++ {
		// i + k 不能超过 数字堆的长度
		for i := 0; i+k < len(piles); i++ {
			j := i + k
			// 如果只有这一堆了
			if i == j {
				dp[i][j] = piles[i]
			} else {
				// 当前堆dp[i][j]的最大差值 为 (取第一堆值d[i] - 从第二堆开始的d[i+1][j])  和 (取最后一堆d[j] - 到倒数第二堆dp[i][j-1]) 中取最大值
				dp[i][j] = int(math.Max(float64(piles[i]-dp[i+1][j]), float64(piles[j]-dp[i][j-1])))
			}
		}
	}
	return dp[0][len(piles)-1] > 0
}
