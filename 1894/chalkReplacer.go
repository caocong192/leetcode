package chalkReplacer

import "sort"

//func chalkReplacer(chalk []int, k int) int {
//	sum := 0
//	for i := 0; i < len(chalk); i++ {
//		sum += chalk[i]
//	}
//
//	k = k % sum
//
//	for i := 0; i < len(chalk); i++ {
//		if k < chalk[i] {
//			return i
//		}
//		k -= chalk[i]
//	}
//	return 0
//}

func chalkReplacer(chalk []int, k int) int {
	if chalk[0] > k {
		return 0
	}

	// 前缀和
	n := len(chalk)
	for i := 1; i < n; i++ {
		chalk[i] += chalk[i-1]
		if chalk[i] > k {
			return i
		}
	}

	k %= chalk[n-1]

	// 二分查找
	return sort.SearchInts(chalk, k+1)
}
