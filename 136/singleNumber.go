package singleNumber

func singleNumber(nums []int) int {
	var m = map[int]int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] == 2 {
			delete(m, nums[i])
		}
	}

	for key, _ := range m {
		return key
	}

	return 0
}
