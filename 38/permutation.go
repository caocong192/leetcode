package permutation

import (
	"sort"
)

func permutation(s string) (ans []string) {
	t := []byte(s)
	// 将字节切片 t 进行从小到大排序, 将相同元素排列在一起
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })

	n := len(t)
	// 保存可能的值
	val := make([]byte, 0, n)
	// 标识该位置元素是否使用
	valid := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		// 当所有字符串使用完了 则添加该元素, 并返回
		if i == n {
			ans = append(ans, string(val))
			return
		}

		for j, b := range valid {
			// 如果 b 为true 表示该字符串使用过了 则continue跳过
			// 如果当前元素 和 上一个元素相同且上一个元素还未填入, 则continue跳过(保证重复元素 一定是从左往右填入) 剪枝
			if b || j > 0 && t[j-1] == t[j] && !valid[j-1] {
				continue
			}
			// 标志为true 表示已使用
			valid[j] = true
			// 追加字符
			val = append(val, t[j])
			// 递归下一个
			dfs(i + 1)
			// 回溯, 还原配置, 剔除当前元素, 并将当前标志位重置为false
			val = val[:len(val)-1]
			valid[j] = false
		}
	}
	dfs(0)
	return
}
