package countBattleships

func countBattleships(board [][]byte) (ans int) {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				if (i == 0 || board[i-1][j] != 'X') && (j == 0 || board[i][j-1] != 'X') {
					ans++
				}
			}
		}
	}

	return ans
}
