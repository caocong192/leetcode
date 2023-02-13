package balancedString

func balancedString(s string) int {

	// 统计每个 Q W E R 的个数
	cnt := map[byte]int{}
	for _, v := range s {
		cnt[byte(v)]++
	}

	// 平均数
	partial := len(s) / 4

	check := func() bool {
		if cnt['Q'] > partial || cnt['W'] > partial || cnt['E'] > partial || cnt['R'] > partial {
			return false
		}
		return true
	}

	// 如果第一次都满足, 则本身是符合要求的字符串,返回0
	if check() {
		return 0
	}

	// res 存储结果, 初始化为s的长度, left,right 分别表示滑动窗口的左和右
	res := len(s)
	right := 0
	for left, v := range s {
		for right < len(s) && !check() {
			cnt[s[right]]--
			right++
		}

		if !check() {
			break
		}

		res = min(res, right-left)
		cnt[byte(v)]++

	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
