package repeatedStringMatch

func repeatedStringMatch(a string, b string) int {

	if flag, ans := kmp(a, b); flag {
		return ans
	}
	return -1
}

func kmp(haystack string, needle string) (bool, int) {
	index := 0
	mainIdx := 0
	subIdx := 0
	next := computeTemporaryArray(needle)
	n := len(needle)/len(haystack) + 3

	for index < len(haystack)*n && subIdx < len(needle) {
		if haystack[mainIdx] == needle[subIdx] {
			index++
			mainIdx = index % len(haystack)
			subIdx++
		} else {
			if subIdx != 0 {
				subIdx = next[subIdx-1]
			} else {
				index++
				mainIdx = index % len(haystack)
			}
		}
	}
	if subIdx == len(needle) {
		if index%len(haystack) == 0 {
			return true, index / len(haystack)
		} else {
			return true, index/len(haystack) + 1
		}
	}
	return false, 0
}

/*
 * computeTemporaryArray: 用于计算子串的 next 数组
 * next: 存储子串的 next 下标
 */
func computeTemporaryArray(needle string) []int {
	next := make([]int, len(needle))
	index := 0
	i := 1
	for i < len(needle) {
		if needle[i] == needle[index] {
			next[i] = index + 1
			i++
			index++
		} else {
			if index != 0 {
				index = next[index-1]
			} else {
				i++
			}
		}
	}
	return next
}
