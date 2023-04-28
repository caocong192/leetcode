package longestStrChain

import "sort"

func longestStrChain(words []string) (ans int) {
	cnt := map[string]int{}

	// 从长度 小到大 排序
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		}
		return false
	})

	for _, word := range words {
		// 初始值1
		cnt[word] = 1
		for i := range word {
			// 遍历每个字符, 拼接获取子字符
			preWord := word[:i] + word[i+1:]
			// 如果子字符存在过, 并且 值+1 大于 当前字符的值
			if cnt[preWord] > 0 && cnt[preWord]+1 > cnt[word] {
				cnt[word] = cnt[preWord] + 1
			}
		}

		if ans < cnt[word] {
			ans = cnt[word]
		}
	}
	return
}
