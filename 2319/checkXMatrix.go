package checkXMatrix

func checkXMatrix(grid [][]int) bool {

	// 获取矩阵长度
	l := len(grid)

	// 遍历矩阵
	for index1, i := range grid {
		for index2, j := range i {
			// 判断是否是对角线 且 值为0
			if isCatercorner(index1, index2, l) && j == 0 {
				return false
			} else if !isCatercorner(index1, index2, l) && j != 0 {
				return false
			}
		}
	}
	return true
}

// 判断是否是对角线
func isCatercorner(n, m, l int) bool {
	if n == m || n+m == l-1 {
		return true
	}
	return false
}
