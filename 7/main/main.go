package main

import "fmt"

func reverse(x int) int {

	ret := 0
	for {
		if x != 0 {
			if ret*10 < -2<<31/10 || ret > (2<<31-1)*10 {
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

	x := -321
	ret := reverse(x)
	fmt.Println(ret)
}
