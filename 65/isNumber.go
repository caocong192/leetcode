package isNumber

import (
	"regexp"
	"strings"
)

func isNumber(s string) bool {
	// 去掉左右空格
	s = strings.TrimSpace(s)
	// 正则方法获取结果
	ret,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))([eE][\\+\\-]?[0-9]+)?$", s)
	return ret
}