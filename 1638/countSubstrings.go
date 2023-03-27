package countSubstrings

func countSubstrings(s string, t string) (ans int) {
	m, n := len(s), len(t)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := 0
			for k := 0; k+i < m && k+j < n; k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff > 1 {
					break
				}
				if diff == 1 {
					ans++
				}
			}
		}
	}

	return
}
