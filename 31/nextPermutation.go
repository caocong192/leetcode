package nextPermutation

func nextPermutation(nums []int) []int {
	/*
		方法一：两遍扫描
		第一遍: 从右往左找, 找到一个较大数 和 较小数, 切较大数在较小数右边。 进行交换
		第二遍: 重新排列右边的数字
	*/

	if len(nums) == 1 {
		return nums
	}

	// 第一遍: 从右往左找, 找到一个较大数 和 较小数, 切较大数在较小数右边。 进行交换
	// 找到 较小的数 索引 i
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; {
		i--
	}
	if i >= 0 {
		// 找到 较大的数 索引 j
		j := len(nums) - 1
		for ; j > 0 && nums[j] <= nums[i]; {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 第二遍: 重新排列 i 右边的数字
	for i = i + 1; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
