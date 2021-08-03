package networkDelayTime

import "math"

func networkDelayTime(times [][]int, n, k int) (ans int) {

	// 取最大int值的一半, 则两个inf累加不会越界
	const inf = math.MaxInt64 / 2

	// 定义一个有向边 graph , 初始化值为 inf
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = inf
		}
	}

	// 将times中的值 赋值给 graph
	for i := range times {
		row, col := times[i][0]-1, times[i][1]-1
		graph[row][col] = times[i][2]
	}

	// 定义 distance 为每个节点与 开始节点k 的距离
	distance := make([]int, n)
	for i := range distance {
		distance[i] = inf
	}

	// k 点初始化距离为0
	distance[k-1] = 0

	// used 表示这个节点是否被使用
	used := make([]bool, n)

	// 取出与 顶点k 距离最小的点
	for i := 0; i < n; i++ {
		// x 表示离k点最近的点
		x := -1
		for y, u := range used {
			// x == -1 确保能一直取到 used 中的第一个点
			if !u && (x == -1 || distance[y] < distance[x]) {
				x = y
			}
		}
		used[x] = true

		// 更新与 x 相关的点的距离
		for y, dis := range graph[x] {
			distance[y] = min(distance[y], dis+distance[x])
		}
	}

	// 取出 distance 中最大的值
	for _, dis := range distance {
		if dis == inf {
			return -1
		}
		ans = max(ans, dis)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
