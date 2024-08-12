package numberOfRightTriangles

func numberOfRightTriangles(grid [][]int) (ans int64) {
	//if len(grid) <= 1 || len(grid[0]) <= 1 {
	//	return
	//}
	col := make([]int, len(grid[0])+1)
	row := make([]int, len(grid)+1)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			row[i] += grid[i][j]
			col[j] += grid[i][j]
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				ans += int64(((row[i]) - 1) * (col[j] - 1))
			}
		}
	}
	return
}
