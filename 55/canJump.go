package canJump

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	step := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+1 >= step {
			step = nums[i]
		} else {
			step--
		}

		if step == 0 && i != len(nums)-1 {
			return false
		}
	}
	return true
}
