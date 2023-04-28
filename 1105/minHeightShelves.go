package minHeightShelves

func minHeightShelves(books [][]int, shelfWidth int) int {

	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1000 * 1000
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		maxHeight, currWidth := 0, 0
		for j := i; j >= 0; j-- {
			currWidth += books[j][0]
			if currWidth > shelfWidth {
				break
			}
			maxHeight = max(maxHeight, books[j][1])
			dp[i+1] = min(dp[i+1], dp[j]+maxHeight)
		}
	}

	return dp[n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
