package strStr

/**
 * Slow method of pattern matching
*/
//func strStr(haystack string, needle string) int {
//	if needle == "" || haystack == needle {
//		return 0
//	}
//
//	n := len(needle)
//	for i := 0; i < len(haystack)-n+1; i++ {
//		if haystack[i:i+n] == needle {
//			return i
//		}
//	}
//
//	return -1
//}

/**
 * Date 22/12/2021
 * @author lucien.cao
 *
 * Do pattern matching using KMP algorithm
 *
 * Runtime complexity - O(m + n) where m is length of text and n is length of pattern
 * Space complexity - O(n)
 */

func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}

	return kmp(haystack, needle)
}

/**
 * KMP algorithm of pattern matching.
 */

func kmp(haystack string, needle string) int {
	mainIdx := 0
	subIdx := 0
	next := computeTemporaryArray(needle)
	for mainIdx < len(haystack) && subIdx < len(needle) {
		if haystack[mainIdx] == needle[subIdx] {
			mainIdx++
			subIdx++
		} else {
			if subIdx != 0 {
				subIdx = next[subIdx-1]
			} else {
				mainIdx++
			}
		}
	}

	if subIdx == len(needle) {
		return mainIdx - subIdx
	} else {
		return -1
	}
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
