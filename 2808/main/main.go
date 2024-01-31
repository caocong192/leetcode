package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4, 5}
	for n := range l {
		fmt.Println(n)
	}
}
