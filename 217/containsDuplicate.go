package containsDuplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] == 2 {
			return true
		}
	}
	return false
}

//func containsDuplicate(nums []int) bool {
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//	return false
//}
