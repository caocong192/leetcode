package maximumNumberOfStringPairs

func maximumNumberOfStringPairs(words []string) (ans int) {
	n := len(words)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if words[i][0] == words[j][1] && words[i][1] == words[j][0] {
				ans++
				break
			}
		}
	}
	return
}
