package canConstruct

func canConstruct(ransomNote string, magazine string) bool {

	// hash表. 统计magazine中字符串个数
	m := map[rune]int{}
	for _, v := range magazine {
		m[v]++
	}

	// 遍历赎金信中的字符串, 从hash表中匹配, 如果没有或者等于0 则不满足,返回flase
	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		} else {
			m[v]--
		}
	}
	return true
}
