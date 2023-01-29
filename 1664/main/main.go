package main

import "fmt"

func waysToMakeFair(nums []int) int {
	var odd, even int

	for i := 0; i < len(nums); i += 2 {
		even += nums[i]
		if i+1 < len(nums) {
			odd += nums[i+1]
		}
	}
	fmt.Printf("odd %d,  even: %d\n", odd, even)

	var ans int

	var oddTemp, evenTemp int

	for i := len(nums) - 1; i >= 0; i-- {
		if i%2 == 0 {
			even -= nums[i]
		} else {
			odd -= nums[i]
		}

		fmt.Printf("odd %d,  even: %d, oddTemp: %d, evenTemp: %d\n", odd, even, oddTemp, evenTemp)
		if odd+oddTemp == even+evenTemp {
			ans ++
		}

		if i%2 == 0 {
			oddTemp += nums[i]
		} else {
			evenTemp += nums[i]
		}

	}

	return ans
}

func main() {
	ans := waysToMakeFair([]int{1,2,3})
	fmt.Println(ans)

}
