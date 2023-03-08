package maxValue

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 初始化 cache
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}

		p := &cache[i][j]
		if *p > 0 {
			return *p
		}

		*p = max(dfs(i-1, j)+grid[i][j], dfs(i, j-1)+grid[i][j])
		return *p
	}
	return dfs(m-1, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
