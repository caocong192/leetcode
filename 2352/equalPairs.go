package equalPairs

import "fmt"

func equalPairs(grid [][]int) (ans int) {
	n := len(grid)
	// hash表. 存储每一行 字符串
	cnt := make(map[string]int)
	for _, row := range grid {
		cnt[fmt.Sprint(row)]++
	}

	for i := 0; i < n; i++ {
		// 遍历每一列
		var arr []int
		for j := 0; j < n; j++ {
			arr = append(arr, grid[j][i])
		}
		if val, ok := cnt[fmt.Sprint(arr)]; ok {
			ans += val
		}
	}
	return
}
