package sortArrayByParity

func sortArrayByParity(nums []int) []int {
	// 双指针
	left, right := 0, len(nums)-1
	for ; left < right; {
		if isOddNumber(nums[left]) && !isOddNumber(nums[right]) {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if !isOddNumber(nums[left]) {
			left++
		} else if isOddNumber(nums[right]) {
			right--
		}
	}

	return nums
}

// isOddNumber 判断是否是奇数
func isOddNumber(num int) bool {
	if num%2 == 1 {
		return true
	}
	return false
}
