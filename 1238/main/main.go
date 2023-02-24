package main

import "fmt"

func main() {
	n := 2
	start := 3
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = (i >> 1) ^ i ^ start
	}

	fmt.Println(1 >> 1)
	fmt.Println(2 >> 1)
	fmt.Println(3 >> 1)
	fmt.Println(4 >> 1)
	fmt.Println(5 >> 1)
	fmt.Println(6 >> 1)
	fmt.Println(7 >> 1)
}
