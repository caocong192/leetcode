package isValid

func isValid(s string) bool {
	// 如果字符串s的长度为基数, 直接返回false
	if len(s) % 2 == 1{
		return false
	}

	// 哈希表
	paris := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	// 栈
	stack := []byte{}

	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 如果是左括号 { [ (, 追加到栈中
		if paris[s[i]] > 0 {
			stack = append(stack, s[i])
		} else {
			// 当右括号时, 栈为空, 或者与最后一个栈不匹配
			if len(stack) == 0 ||  s[i] != paris[stack[len(stack)-1]] {
				return false
			}
			// 最后一个元素移除栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
