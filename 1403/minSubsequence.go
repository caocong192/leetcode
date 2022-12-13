package minSubsequence

import (
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func minSubsequence(nums []int) []int {
	sort.Sort(IntSlice(nums))

	total := 0
	for _, num := range nums {
		total += num
	}

	for i, sum := 0, 0; ; i++ {
		sum += nums[i]
		total -= nums[i]

		if sum > total {
			return nums[:i+1]
		}
	}

}
