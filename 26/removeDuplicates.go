package removeDuplicates

func removeDuplicates(nums []int) int {
	// 如果数组为空, 直接返回0
	if len(nums) == 0 {
		return 0
	}

	// 双指针, 如果i下标元素不与i-1元素相同, 则替换j下标元素
	fast, slow := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
