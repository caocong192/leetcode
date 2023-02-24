package searchInsert

func searchInsert(nums []int, target int) int {
	return search(nums, target)
}

// search 二分法查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		num := nums[mid]
		if num > target {
			right = mid
		} else if num < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
