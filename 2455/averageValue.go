package averageValue

func averageValue(nums []int) (ans int) {
	var num = 0
	for _, v := range nums {
		if v%3 == 0 && v%2 == 0 {
			ans += v
			num += 1
		}
	}
	if num == 0 {
		return num
	}
	return ans / num
}
