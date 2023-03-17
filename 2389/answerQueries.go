package answerQueries

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	var ans []int
	for _, q := range queries {
		sum := 0
		for j := 0; j < len(nums); j++ {
			if q >= nums[j] {
				q -= nums[j]
				sum++
				continue
			}
			break
		}
		ans = append(ans, sum)
	}

	return ans
}
