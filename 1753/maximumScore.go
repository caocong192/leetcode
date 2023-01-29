package maximumScore

func maximumScore(a int, b int, c int) int {
	sum := a + b + c
	maxNum := max(a, max(b, c))

	// 如果最大值 maxNum c 大于 其它两个值a,b, 则返回a+b
	if maxNum > sum/2 {
		return sum - maxNum
	} else {
		return sum / 2
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//func maximumScore(a int, b int, c int) (want int) {
//	for true {
//		if a == 0 && b == 0 || a == 0 && c == 0 || b == 0 && c == 0 {
//			return
//		}
//
//		a, b, c = GetMaxTwo(a, b, c)
//		want += 1
//	}
//	return
//}
//
//// GetMaxTwo 获取最大的两个数 进行-1操作
//func GetMaxTwo(a, b, c int) (d, e, f int) {
//	if a <= b && a <= c {
//		b -= 1
//		c -= 1
//	} else if b <= a && b <= c {
//		a -= 1
//		c -= 1
//	} else if c <= a && c <= b {
//		a -= 1
//		b -= 1
//	}
//	return a, b, c
//}
