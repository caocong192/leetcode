package main

import "fmt"

func isPalindrome(x int) bool {
	// 反转一半的数字
	// 如果是负数, 直接返回false
	if x < 0 || x%10 == 0 {
		return false
	}

	// 定义反转后的正数 reNum
	reNum := 0

	// 反转一半整数逻辑
	for reNum < x {
		reNum = reNum*10 + x%10
		x /= 10
		fmt.Println(x, reNum)
	}

	// 如果反转后和当前x相同, 或者反转后/10 和x的结果相同 则为true
	if reNum == x || reNum/10 == x {
		return true
	}
	return false
}

func main() {
	x := 10
	ret := isPalindrome(x)
	fmt.Println(ret)
}
