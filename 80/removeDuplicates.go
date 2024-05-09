package removeDuplicates

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == 0 || i == 1 {
			continue
		}
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
