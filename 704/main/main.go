package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	index := search(nums, target)
	fmt.Println(index)

}


func search(nums []int, target int) int {
	// 第一个数大于target 或者 最后一个数小于 target
	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	start := 0
	end := len(nums)

	for {
		if nums[(end+start)/2] == target {
			return (end + start) / 2
		}
		fmt.Println(start, end)

		if end-start == 1 {
			break
		}

		if nums[(end+start)/2] > target {
			end = (end + start) / 2
		} else {
			start = (end + start) / 2
		}
	}
	return -1
}