package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	nums := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	sort.Sort(IntSlice(nums))

	total := 0
	for _, num := range nums {
		total += num
	}

	for i, sum := 0, 0; ; i++ {
		sum += nums[i]
		total -= nums[i]

		if sum > total {
			fmt.Println(nums)
			fmt.Println(sum)
			fmt.Println(nums[:i+1])
			return
		}
	}
}
