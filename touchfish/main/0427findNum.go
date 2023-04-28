package main

import "fmt"

/*
5个一位整数之和为30
其中一个是1，一个是8
而这5个数的乘积是2520。
你能说出余下的是哪3个数吗?
*/

func main() {
	var a, b = 1, 8
	var leftNum = []int{0, 2, 3, 4, 5, 6, 7, 9}
	for _, i := range leftNum {
		for _, j := range leftNum {
			for _, k := range leftNum {
				if i == j || i == k || j == k {
					continue
				}
				if i+j+k == 21 && i*j*k*8 == 2520 {
					fmt.Printf("这五个数分别是: %d,%d,%d,%d,%d\n", a, b, i, j, k)
					return
				}
			}
		}
	}
}
