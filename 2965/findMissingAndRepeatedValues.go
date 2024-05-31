package findMissingAndRepeatedValues

func findMissingAndRepeatedValues(grid [][]int) []int {
	length := len(grid)
	b := (1 + length*length) * length * length / 2
	m := map[int]int{}

	a := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			m[grid[i][j]]++

			// 如果  出现两次 则为a
			if m[grid[i][j]] > 1 {
				a = grid[i][j]
				continue
			}

			b -= grid[i][j]
		}
	}
	return []int{a, b}
}
