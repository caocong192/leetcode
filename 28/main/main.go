package main

import "fmt"

func computeTemporaryArray(needle string) []int {
	next := make([]int, len(needle))
	index := 0
	i := 1
	for i < len(needle) {
		if needle[i] == needle[index] {
			next[i] = index + 1
			i++
			index++
		} else {
			if index != 0 {
				index = next[index-1]
			} else {
				i++
			}
		}
	}
	return next
}

func main() {
	needle := "abcdabc"
	ret := computeTemporaryArray(needle)
	for i, v := range ret {
		fmt.Println(i, v)
	}
}
