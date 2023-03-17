package getRow

func getRow(rowIndex int) []int {

	var ans [][]int
	for i := 0; i < rowIndex; i++ {
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
	return ans[rowIndex-1]
}
