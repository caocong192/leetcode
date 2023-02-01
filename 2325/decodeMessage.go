package decodeMessage

func decodeMessage(key string, message string) string {
	curr := byte('a')

	m := make(map[rune]byte)

	for _, v := range key {
		if v != ' ' && m[v] == 0 {
			m[v] = curr
			curr++
		}
	}

	res := []byte(message)

	for i, v := range message {
		if v != ' ' {
			res[i] = m[v]
		}
	}

	return string(res)
}

//func decodeMessage(key string, message string) string {
//
//	// m 用map来存储每个字符
//	m := make(map[string]string)
//	// 从小写字母a开始
//	x := 97
//	// res map用于存储解密对照表
//	res := make(map[string]string)
//
//	for _, v := range key {
//		// 如果是空字符串 或者 已经统计过的 字符 跳过
//		if string(v) == " " || m[string(v)] != "" {
//			continue
//		}
//		// 存入临时字符map m
//		m[string(v)] = string(v)
//		// 存入 解密对照表 res
//		//res[string(v)] = string(x)
//		res[string(v)] = fmt.Sprintf("%c", x)
//		// a-z 往后走
//		x++
//	}
//
//	// 解密message
//	var want string
//	for _, v := range message {
//		if string(v) == " " {
//			want += " "
//		} else {
//			want += res[string(v)]
//		}
//	}
//
//	return want
//}
