package maxProfit

func maxProfit(prices []int) (ans int) {
	var min = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > ans {
			ans = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return ans
}
