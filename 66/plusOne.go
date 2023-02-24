package plusOne

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		if digits[i]+1 <= 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
	}
	if i == -1 {
		ans := make([]int, len(digits)+1)
		ans[0] = 1
		return ans
	}
	return digits
}
