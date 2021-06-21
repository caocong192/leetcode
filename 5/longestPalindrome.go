package longestPalindrome

func longestPalindrome(s string) string {
	// start, end 保存最大回文的下标
	start, end := 0, 0
	for i := 0; i < len(s); i++ { // s := "12321"
		// 以 s[i] 为中心算最大回文串, left1,right1 分别判断左右边界
		left1, right1 := expandPalindrome(s, i, i)
		// 以 s[i, i+1] 为中心算最大回文串, left2,right2 分别判断左右边界
		left2, right2 := expandPalindrome(s, i, i+1)

		// 保存最大的回文逻辑
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandPalindrome(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
