package rotate

func rotate(nums []int, k int) {
	k = k % len(nums)
	newNums := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, newNums)
}
