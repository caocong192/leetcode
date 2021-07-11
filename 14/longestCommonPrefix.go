package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	// 如果 strs 中只有一个字符串
	if len(strs) == 1 {
		return strs[0]
	}

	var ans []byte

	// 取第一个字符串为判断基准
	for i := 0; i < len(strs[0]); i++ {
		// 每次追加一个字符
		ans = append(ans, strs[0][i])
		// 遍历所有字符串
		for _, val := range strs {
			// 如果长度不够了， 直接返回结果
			if len(val) < i+1 {
				return string(ans)[0 : len(ans)-1]
			}
			// 判断是否有相同的前缀, 如果是 继续匹配下一个
			if string(ans) == val[0:i+1] {
				continue
			}
			// 否者 返回结果
			return string(ans)[0 : len(ans)-1]
		}
	}

	// 如果都是空字符串, 则前缀为空, ans长度为0, 返回空
	if len(ans) == 0 {
		return ""
	}
	return string(ans)
}
