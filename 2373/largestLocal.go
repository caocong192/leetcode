package largestLocal

func largestLocal(grid [][]int) [][]int {
	n := len(grid[0])
	var ans [][]int

	for i := 0; i < n-2; i++ {
		var a []int
		for j := 0; j < n-2; j++ {
			a = append(a, getMaxValue(grid, i, j))
		}
		ans = append(ans, a)
	}
	return ans
}

// getMaxValue 获取 3X3 矩阵中最大的值
func getMaxValue(grid [][]int, i, j int) (max int) {
	var step = 2
	for ; step >= 0; step-- {
		var step2 = 2
		for ; step2 >= 0; step2-- {
			if grid[i+step][j+step2] > max {
				max = grid[i+step][j+step2]
			}
		}
	}
	return
}
