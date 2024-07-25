package minimumOperations

func minimumOperations(num string) int {
	// 00, 25, 50, 75
	n := len(num) - 1
	find0, find5 := false, false
	for i := n; i >= 0; i-- {
		if num[i] == '0' {
			if find0 {
				return n - i - 1
			}
			find0 = true
		} else if num[i] == '5' {
			if find0 {
				return n - i - 1
			}
			find5 = true
		} else if num[i] == '2' {
			if find5 {
				return n - i - 1
			}
		} else if num[i] == '7' {
			if find5 {
				return n - i - 1
			}
		}
	}
	if find0 {
		return n
	}
	return n + 1
}
