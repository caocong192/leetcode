package countSegments

import (
	"strings"
)

func countSegments(s string) (ans int) {

	//// 删除前置的空格
	//for true {
	//	if len(s) != 0 && s[0] == ' ' {
	//		s = s[1:]
	//		continue
	//	}
	//	break
	//}
	//
	//// 删除后置的空格
	//for true {
	//	if len(s) != 0 && s[len(s)-1] == ' ' {
	//		s = s[:len(s)-1]
	//		continue
	//	}
	//	break
	//}
	//
	//// 如果是空字符串, 则反回0
	//if len(s) == 0 {
	//	return
	//}

	// 按空格切分
	l := strings.Split(s, " ")
	for i := 0; i < len(l); i++ {
		if len(l[i]) != 0 {
			ans++
		}
	}
	return
}
