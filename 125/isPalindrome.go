package isPalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	// 转小写字母
	s = strings.ToLower(s)

	// 遍历字符串s, 去除特殊字符, 存入新的字符串 newStr
	var newStr []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			newStr = append(newStr, s[i])
		}
	}

	// 遍历 newStr, 从两端判断字符是否相等
	for i := 0; i < len(newStr)/2; i++ {
		if newStr[i] != newStr[len(newStr)-i-1] {
			return false
		}
	}
	return true
}
