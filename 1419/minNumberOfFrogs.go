package minNumberOfFrogs

func minNumberOfFrogs(croakOfFrogs string) (ans int) {
	m := map[rune]int{
		'c': 0,
		'r': 0,
		'o': 0,
		'a': 0,
		'k': 0,
	}

	for i := 0; i < len(croakOfFrogs); i++ {
		m[rune(croakOfFrogs[i])]++

		if m['c'] >= m['r'] && m['r'] >= m['o'] && m['o'] >= m['a'] && m['a'] >= m['k'] {
			ans = max(ans, m['c']-m['k'])
		} else {
			return -1
		}
	}
	if m['c'] == m['r'] && m['r'] == m['o'] && m['o'] == m['a'] && m['a'] == m['k'] {
		return
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
