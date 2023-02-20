package largest1BorderedSquare

func largest1BorderedSquare(grid [][]int) int {
	// [i][j] 表示 以 i,j 为底点,
	// left[i][j] 左边 都为 1 最长的边
	// up[i][j] 上边 都为 1 最长的边
	n := len(grid)
	m := len(grid[0])
	left := make([][]int, n+1)
	up := make([][]int, n+1)
	for i := range left {
		left[i] = make([]int, m+1)
		up[i] = make([]int, m+1)
	}

	maxBroder := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if grid[i-1][j-1] == 1 {
				left[i][j] = left[i][j-1] + 1
				up[i][j] = up[i-1][j] + 1
				broder := min(left[i][j], up[i][j])

				// 如果当前边 小于 之前得到的最大边, 直接跳过后面的计算
				if broder <= maxBroder {
					continue
				}

				// 判断 另外 两个边 是否大于 broder
				for left[i-broder+1][j] < broder || up[i][j-broder+1] < broder {
					broder--
				}
				maxBroder = max(maxBroder, broder)
			}
		}
	}
	return maxBroder * maxBroder
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
