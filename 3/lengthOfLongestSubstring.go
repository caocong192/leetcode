package lengthOfLongestSubstring

import "math"

/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {

	// 记录每个字符是否出现
	m := map[byte]int{}
	// sLen 字符串长度
	sLen := len(s)

	// right 为滑动窗口的右侧, ret保存每次滑动的最大不重复字符串的长度值
	right, ret := -1, 0.0

	// 遍历字符串
	for left := 0; left < sLen; left++ {
		// 从第二次开始, 每次删除前一个字符
		if left != 0 {
			delete(m, s[left-1])
		}

		// 获取最大不重复字符串逻辑
		for right+1 < sLen && m[s[right+1]] == 0 {
			m[s[right+1]] = 1
			right++
		}
		ret = math.Max(ret, float64(right-left)+1)
	}
	return int(ret)
}
