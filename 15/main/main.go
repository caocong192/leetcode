package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var num uint32
	num = 11111111111111111111111111111101
	fmt.Println(num)
	ans := hammingWeight(num)
	fmt.Println(ans)

}

func hammingWeight(num uint32) int {
	ans := bits.OnesCount32(num)
	return ans
}
