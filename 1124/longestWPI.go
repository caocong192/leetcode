package longestWPI

func longestWPI(hours []int) int {
	// 后缀和 s
	n := len(hours)
	s := make([]int, n)

	// 如果 hours[i] 大于 8 则值改为1, 否者改为-1
	// s[i] 代表 从 下标 i 到 n 的和
	if hours[n-1] > 8 {
		hours[n-1] = 1
		s[n-1] = 1
	} else {
		hours[n-1] = -1
		s[n-1] = -1
	}

	for i := len(hours) - 2; i >= 0; i-- {
		if hours[i] > 8 {
			hours[i] = 1
			s[i] = s[i+1] + 1
		} else {
			hours[i] = -1
			s[i] = s[i+1] - 1
		}
	}

	ans := 0
	for i, _ := range hours {
		if s[i] > 0 {
			// 如果 s[i] 大于 0, 则最大长度就是 i-n, 直接和 ans比较, 返回最大值
			ans = max(ans, n-i)
			break
		}

		curr := s[i]
		// 倒序遍历, 开始逐步 减去 hours[j], 直到 curr > 0, 则从 i 开始的最大长度就是 i - j
		for j := n - 1; j >= 0; j-- {
			curr -= hours[j]
			if curr > 0 {
				ans = max(ans, j-i)
				break
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
