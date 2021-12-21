package removeElement

//func removeElement(nums []int, val int) int {
//	n := len(nums)
//
//	for i := 0; i < n; {
//		if nums[i] == val {
//			nums[i], nums[n-1] = nums[n-1], nums[i]
//			n--
//			continue
//		}
//		i++
//	}
//	return n
//}

//func removeElement(nums []int, val int) int {
//	left := 0
//	for _, v := range nums {
//		if v != val {
//			nums[left] = v
//			left++
//		}
//	}
//	return left
//}

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return right
}
