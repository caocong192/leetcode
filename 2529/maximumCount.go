package maximumCount

func maximumCount(nums []int) (ans int) {
	n := len(nums)
	for i, num := range nums {
		if num < 0 {
			ans++
		}

		if num > 0 {
			if n-i > ans {
				ans = n - i
				return
			}
		}
	}
	return
}
