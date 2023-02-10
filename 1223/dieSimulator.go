package dieSimulator

func dieSimulator(n int, rollMax []int) (ans int) {
	const mod int = 1e9 + 7

	// 记忆化搜索
	const m = 6
	// 存储 n int, last int, left int, 次数,上一次数的值, 剩余次数的状态
	cache := make([][m][]int, n)
	for i := range cache {
		for j := range cache[i] {
			cache[i][j] = make([]int, rollMax[j])
			for k := range cache[i][j] {
				// 初始化标记为 -1
				cache[i][j][k] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(n int, last int, left int) (res int) {
		if n == 0 {
			return 1
		}

		// 如果之前查到过, 则使用缓存
		c := &cache[n][last][left]
		if *c != -1 {
			return *c
		}

		for j, max := range rollMax {
			if j != last {
				res += dfs(n-1, j, max-1)
			} else if left > 0 {
				res += dfs(n-1, j, left-1)
			}
		}
		*c = res % mod
		return res % mod
	}

	for j, max := range rollMax {
		ans += dfs(n-1, j, max-1)
	}

	return ans % mod
}
