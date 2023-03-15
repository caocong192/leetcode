package generate

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	var ans [][]int
	for i := 0; i < numRows; i++ {
		var a []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				a = append(a, 1)
			} else {
				a = append(a, ans[i-1][j-1]+ans[i-1][j])
			}
		}
		ans = append(ans, a)
	}
	return ans
}
