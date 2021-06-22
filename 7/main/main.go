package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	ret := 0
	for {
		if x != 0 {
			if ret < -math.MaxInt32/10 || ret > math.MaxInt32/10 {
				return 0
			}
			digit := x % 10
			x = x / 10
			ret = ret*10 + digit
			fmt.Printf("ret: %d\n", ret)
		} else {
			return ret
		}
	}
}

func main() {
	fmt.Println(1<<31)

	x := -1463847412
	ret := reverse(x)
	fmt.Println(ret)
}
