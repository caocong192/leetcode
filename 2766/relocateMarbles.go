package relocateMarbles

import "sort"

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) (ans []int) {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	for i := 0; i < len(moveFrom); i++ {
		m[moveFrom[i]] = false
		m[moveTo[i]] = true
	}

	for key, value := range m {
		if value == true {
			ans = append(ans, key)
		}
	}
	sort.Ints(ans)
	return
}
