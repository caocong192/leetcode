package longestDecomposition

import "fmt"

func longestDecomposition(text string) (ans int) {
	var l, r = 0, len(text)

	var num = 0
	for l < r/2 {

		if text[l-num:l+1] == text[r-l-1:r-l+num] {
			fmt.Println(text[l-num : l+1])
			fmt.Println(text[r-l-1 : r-l+num])
			ans++
			num = 0
		} else {
			num++
		}
		l++
	}

	// 如果中间还有一位数 或者 一段数字, 加1
	if 2*l < r || num != 0 {
		return ans*2 + 1
	} else {
		return ans * 2
	}
}
