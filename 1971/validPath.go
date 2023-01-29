package validPath

func validPath(n int, edges [][]int, source int, destination int) bool {
	valid := make([]bool, n)
	g := make([][]int, n)

	for _, v := range edges {
		a, b := v[0], v[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	// dfs 深度优先搜索
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == destination {
			return true
		}

		valid[i] = true

		for _, j := range g[i] {
			if !valid[j] && dfs(j) {
				return true
			}
		}
		return false
	}
	return dfs(source)
}


