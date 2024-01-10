package minLength

//import "strings"
//
//func minLength(s string) int {
//	match1 := "AB"
//	match2 := "CD"
//	for strings.Contains(s, match1) || strings.Contains(s, match2) {
//		s = strings.Replace(s, match1, "", -1)
//		s = strings.Replace(s, match2, "", -1)
//	}
//
//	return len(s)
//}

func minLength(s string) int {
	// 栈, 栈顶元素为AB or CD 弹出
	var stack []byte
	for i := range s {
		stack = append(stack, s[i])
		l := len(stack)
		for l >= 2 && (string(stack[l-2:]) == "AB" || string(stack[l-2:]) == "CD") {
			stack = stack[:l-2]
		}
	}
	return len(stack)
}
