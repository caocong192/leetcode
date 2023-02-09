package main

import (
	"fmt"
	"github.com/leetcode/leetcode/1797/AuthenticationManager"
)

func main() {
	// AuthenticationManager Constructor
	authCode := AuthenticationManager.Constructor(5)

	// AuthenticationManager renew
	authCode.Renew("aaa", 1)

	// AuthenticationManager generate
	authCode.Generate("aaa", 2)

	// AuthenticationManager countUnexpiredTokens
	count := authCode.CountUnexpiredTokens(6)
	fmt.Printf("当前时间%d 存在的验证码个数: %d\n", 6, count)

	// AuthenticationManager enerate
	authCode.Generate("bbb", 7)

	authCode.Renew("aaa",8)

	authCode.Renew("bbb",10)

	count = authCode.CountUnexpiredTokens(15)
	fmt.Printf("当前时间%d 存在的验证码个数: %d\n", 15, count)
}
