package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {

	// 去除前后空格
	s = strings.TrimSpace(s)

	// 如果第一个字符是字母, 则返回0

	if unicode.IsLower(rune(s[0])) || s[0] == '.' {
		return 0
	}

	flag := 1
	if rune(s[0]) == '-' {
		flag = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsLower(rune(s[i])) || s[i] == '.' || s[i] == ' ' {
			return ret * flag
		}

		val, _ := strconv.Atoi(string(s[i]))
		res := ret*10 + val

		if res*flag >= math.MaxInt32 {
			return math.MaxInt32
		} else if res*flag < -math.MaxInt32 {
			return -math.MaxInt32 - 1
		}
		ret = res
	}
	return ret * flag
}

func main() {
	s := "-91283472332"
	ret := myAtoi(s)
	fmt.Println(ret)

}
