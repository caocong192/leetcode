package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		nums = append(nums[len(nums)-1:], nums[:len(nums)-1]...)
	}

	fmt.Printf("%v", nums)
}
