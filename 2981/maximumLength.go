package maximumLength

func maximumLength(s string) int {
	ans := -1
	length := len(s)

	chs := make([][]int, 26)
	cnt := 0

	for i := 0; i < length; i++ {
		cnt++
		if i+1 == length || s[i] != s[i+1] {
			ch := s[i] - 'a'
			chs[ch] = append(chs[ch], cnt)
			cnt = 0

			for j := len(chs[ch]) - 1; j > 0; j-- {
				if chs[ch][j] > chs[ch][j-1] {
					tmp := chs[ch][j-1]
					chs[ch][j-1] = chs[ch][j]
					chs[ch][j] = tmp
				}
			}

			if len(chs[ch]) > 3 {
				chs[ch] = chs[ch][:len(chs[ch])-1]
			}
		}
	}

	for i := 0; i < 26; i++ {
		if len(chs[i]) > 0 && chs[i][0] > 2 {
			ans = max(ans, chs[i][0]-2)
		}
		if len(chs[i]) > 1 && chs[i][0] > 1 {
			ans = max(ans, min(chs[i][0]-1, chs[i][1]))
		}
		if len(chs[i]) > 2 {
			ans = max(ans, chs[i][2])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
