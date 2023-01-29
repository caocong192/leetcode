package findGCD

import "sort"

func findGCD(nums []int) int {
	sort.Ints(nums)

	maxNum := nums[len(nums)-1]
	minNum := nums[0]

	// 如果最大数能整除最小数, 则返回最小数
	if maxNum%minNum == 0 {
		return minNum
	}

	//for i := minNum / 2; i >= 1; i-- {
	//	if minNum%i == 0 && maxNum%i == 0 {
	//		return i
	//	}
	//}
	return gcd(maxNum, minNum)
}

// gcd 辗转相除法
func gcd(a, b int) int {
	// 如果分子大于分母，需要交换顺序
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
