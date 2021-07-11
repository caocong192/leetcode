package numWays

func numWays(n int, relation [][]int, k int) (ans int) {
	// 构造二位数组, 存放当前位置有多少条路选择
	edges := make([][]int, n)
	// 遍历路径
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}

	// 构造dfs函数
	var dfs func(int, int)
	dfs = func(m int, step int) {
		if step == k {
			if m == n-1 {
				ans++
			}
			return
		}

		for _, next := range edges[m] {
			dfs(next, step+1)
		}
	}
	dfs(0, 0)
	return
}
