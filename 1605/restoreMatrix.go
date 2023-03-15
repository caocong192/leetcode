package restoreMatrix

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	for i, j := 0, 0; i < n && j < m; {
		v := min(rowSum[i], colSum[j])
		matrix[i][j] = v
		rowSum[i] -= v
		colSum[j] -= v
		if rowSum[i] == 0 {
			i++
		}

		if colSum[j] == 0 {
			j++
		}
	}

	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
