package bestHand

//"Flush"：同花，五张相同花色的扑克牌。
//"Three of a Kind"：三条，有 3 张大小相同的扑克牌。
//"Pair"：对子，两张大小一样的扑克牌。
//"High Card"：高牌，五张大小互不相同的扑克牌。

func bestHand(ranks []int, suits []byte) string {

	// Flush
	if suits[0] == suits[1] && suits[0] == suits[2] && suits[0] == suits[3] && suits[0] == suits[4] {
		return "Flush"
	}

	// Three of a kind
	var m = make(map[int]int)
	for _, v := range ranks {
		m[v] += 1
	}
	for _, v := range m {
		if v >= 3 {
			return "Three of a Kind"
		}
	}

	// Pair
	for _, v := range m {
		if v == 2 {
			return "Pair"
		}
	}

	return "High Card"
}
