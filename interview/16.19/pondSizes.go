package pondSizes

import "sort"

func pondSizes(land [][]int) []int {
	var res []int
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j] == 0 {
				res = append(res, pool(land, i, j, 0))
			}
		}
	}

	sort.Ints(res)
	return res
}

func pool(land [][]int, i int, j int, s int) int {
	if i < 0 || j < 0 || i >= len(land) || j >= len(land[i]) || land[i][j] != 0 {
		return s
	}

	if land[i][j] == 0 {
		s += 1
		land[i][j] = -1
	}

	for ii := i - 1; ii <= i+1; ii++ {
		for jj := j - 1; jj <= j+1; jj++ {
			s += pool(land, ii, jj, 0)
		}
	}
	return s
}
