package getLucky

import (
	"fmt"
	"strconv"
)

var m = map[string]string{
	"a": "1",
	"b": "2",
	"c": "3",
	"d": "4",
	"e": "5",
	"f": "6",
	"g": "7",
	"h": "8",
	"i": "9",
	"j": "10",
	"k": "11",
	"l": "12",
	"m": "13",
	"n": "14",
	"o": "15",
	"p": "16",
	"q": "17",
	"r": "18",
	"s": "19",
	"t": "20",
	"u": "21",
	"v": "22",
	"w": "23",
	"x": "24",
	"y": "25",
	"z": "26",
}

func getLucky(s string, k int) int {
	transNum := transform(s)

	for i := 1; i <= k; i++ {
		transNum = countNums(transNum)
	}
	res, _ := strconv.Atoi(transNum)
	return res
}

// 字母转化为数字
func transform(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		res += m[string(s[i])]
	}
	return res
}

// 累计每个数字
func countNums(nums string) string {
	var res int
	for i := 0; i < len(nums); i++ {
		num, _ := strconv.Atoi(string(nums[i]))
		res += num
	}
	return fmt.Sprintf("%v", res)
}
