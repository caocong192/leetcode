package myAtoi

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {

	// 去除前后空格
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// 如果第一个字符是字母或者".", 则返回0
	//if unicode.IsLower(rune(s[0])) || s[0] == '.' {
	//	return 0
	//}

	// flag标识位, 用于保存数字符号
	flag := 1
	if rune(s[0]) == '-' {
		flag = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	// ret保存转换后的数字
	ret := 0
	for i := 0; i < len(s); i++ {
		// 如果该字符是数字或者"." 直接返回
		if unicode.IsLower(rune(s[i])) || s[i] == '.' || s[i] == ' ' || s[i] == '-' || s[i] == '+' {
			return ret * flag
		}

		val, _ := strconv.Atoi(string(s[i]))
		res := ret*10 + val

		// 判断是否越界  范围: [-2^31, 2^31-1]
		if res*flag >= math.MaxInt32 {
			return math.MaxInt32
		} else if res*flag < -math.MaxInt32 {
			return -math.MaxInt32 - 1
		}
		ret = res
	}
	return ret * flag
}
