package maxLength

func maxLength(arr []string) int {
	// 存储所有有效的字符串的切片
	var validArr []string

	// 遍历字符串
	for i := 0; i < len(arr); i++ {
		// 判断当前字符串是否有效, 如果无效, 则跳过后续步骤继续for循环
		if !isValid(arr[i]) {
			continue
		}

		// 遍历所有的有效字符串切片
		for _, validStr := range validArr {
			// 逐个进行拼接
			jointStr := validStr + arr[i]
			// 如果有效, 则加入有效字符串切片中
			if isValid(jointStr) {
				validArr = append(validArr, jointStr)
			}
		}
		// 当前字符串有效, 加入到有效字符串切片中
		validArr = append(validArr, arr[i])
	}

	// 遍历所有的有效字符串切片, 取长度最大的值
	ret := 0
	for _, a := range validArr {
		if ret < len(a) {
			ret = len(a)
		}
	}
	return ret
}

func isValid(s string) bool {
	m := map[byte]int{}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return false
		}
		m[s[i]] = 1
	}
	return true
}
