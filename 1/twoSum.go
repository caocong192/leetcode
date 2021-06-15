package twoSum

func twoSum(nums []int, target int) []int {

	var nums2 = map[int]int{}

	for index1, num1 := range nums {
		nums2[index1] = target - num1

		for index2, num2 := range nums2 {
			if num1 == num2 && index1 != index2 {
				var res = []int{index2, index1}
				return res
			}
		}
	}
	return []int{}
}
