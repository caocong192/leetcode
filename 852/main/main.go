package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	var index int
	var arrLen int
	preIndex := 0

	for {
		arrLen = len(arr)
		index = arrLen / 2
		if arr[index] < arr[index+1] {
			arr = arr[index-1:]
			preIndex += index - 1
			fmt.Println(arr)
		} else if arr[index] < arr[index-1] {
			arr = arr[:index+1]
			fmt.Println(arr)
		} else {
			return index + preIndex
		}
	}
	return 0
}

func main() {

	arr := []int{0, 10, 11, 12, 13, 14, 15, 2}
	fmt.Println(peakIndexInMountainArray(arr))
}
