package waysToMakeFair

func waysToMakeFair(nums []int) (ans int) {
	var odd, even int

	// 先统计奇偶数的和
	for i := 0; i < len(nums); i += 2 {
		even += nums[i]
		if i+1 < len(nums) {
			odd += nums[i+1]
		}
	}

	// 从后面统计奇偶数的和
	var oddTemp, evenTemp int

	// 从数组末尾开始, odd even 一个个减去对应的值
	for i := len(nums) - 1; i >= 0; i-- {
		if i%2 == 0 {
			even -= nums[i]
		} else {
			odd -= nums[i]
		}

		// 判断是否相等
		if odd+oddTemp == even+evenTemp {
			ans ++
		}

		// 这个数必定从偶数变基数, 基数变偶数, 用 oddTemp, evenTemp 存起来
		if i%2 == 0 {
			oddTemp += nums[i]
		} else {
			evenTemp += nums[i]
		}

	}

	return ans
}