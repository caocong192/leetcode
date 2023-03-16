package main

import "fmt"

func main() {
	num := []int{3, 2, 1, 4, 5}
	countSubarrays(num, 4)

}

func countSubarrays(nums []int, k int) int {

	var prefixSum []int

	// 替换数组
	var sum int
	for _, v := range nums {
		if v < k {
			sum -= 1
		} else if v > k {
			sum += 1
		}
		prefixSum = append(prefixSum, sum)
	}

	fmt.Println(prefixSum)

	return 0
}
