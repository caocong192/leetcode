package smallestRepunitDivByK

func smallestRepunitDivByK(k int) int {
	if k == 2 || k == 5 {
		return -1
	}

	cur := 0
	for res := 1; res <= 1000; res++ {
		cur = (10*cur + 1) % k
		if cur == 0 {
			return res
		}

	}
	return -1
}
