package diagonalSort

import "sort"

func diagonalSort(mat [][]int) [][]int {
	n := len(mat)
	m := len(mat[0])
	diag := make([][]int, n+m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			diag[i-j+m] = append(diag[i-j+m], mat[i][j])
		}
	}

	for _, d := range diag {
		sort.Sort(sort.Reverse(sort.IntSlice(d)))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mat[i][j] = diag[i-j+m][len(diag[i-j+m])-1]
			diag[i-j+m] = diag[i-j+m][:len(diag[i-j+m])-1]
		}
	}

	return mat
}
