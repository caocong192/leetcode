package snakesAndLadders

func snakesAndLadders(board [][]int) (ans int) {
	// row 行大小
	row := len(board)
	// 记录bfs中 所有的可能性 队列
	queue := make([]int, 1, 128)
	// ans 记录步数
	ans = 0
	// size 总格子数大小
	size := row * row
	// visited 记录该位置是否走过
	visited := make([]bool, size+1)

	// 从第一个点出发
	queue[0] = 1
	visited[0] = true
	// 当队列没有计算完, 则一直继续
	for len(queue) > 0 {
		// 遍历所有队列
		n := len(queue)
		for i := 0; i < n; i++ {
			// 取出第一个
			p := queue[0]
			queue = queue[1:]

			// 如果到达目的地
			if p == size {
				return
			}

			// 模拟投掷筛子, 1-6
			for j := p + 1; j <= p+6 && j <= size; j++ {
				r, c := getRC(j, row)
				// 跳到指定的地方
				next := j
				if board[r][c] > 0 {
					next = board[r][c]
				}
				// 将下一个点加入队列
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		ans++
	}
	return -1
}

func getRC(j, row int) (r, c int) {
	r = row - (j-1)/row - 1
	// 如果是奇数行
	c = (j - 1) % row
	// 如果是偶数行
	if (row-r)%2 == 0 {
		c = row - 1 - c
	}
	return
}

