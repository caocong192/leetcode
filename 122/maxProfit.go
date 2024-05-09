package maxProfit

func maxProfit(prices []int) (ans int) {
	lowPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < lowPrice {
			lowPrice = prices[i]
		} else {
			ans += prices[i] - lowPrice
			lowPrice = prices[i]
		}
	}
	return
}
