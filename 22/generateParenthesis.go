package generateParenthesis

func generateParenthesis(n int) (ans []string) {
	// 构造dfs函数
	var dfs func(string, int, int, int)
	dfs = func(str string, l int, r int, n int) {
		// 剪枝, 左右括号大于 n时 或者右括号大于左括号时 退出
		if l > n || r > n || r > l {
			return
		}
		if l == n && r == n {
			ans = append(ans, str)
			return
		}
		dfs(str+"(", l+1, r, n)
		dfs(str+")", l, r+1, n)
	}

	dfs("", 0, 0, n)
	return ans
}
