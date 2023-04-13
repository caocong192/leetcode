package mostFrequentEven

func mostFrequentEven(nums []int) (ans int) {
	var m = map[int]int{}

	for _, num := range nums {
		if num%2 == 0 {
			m[num]++
		}
	}

	if len(m) == 0 {
		return -1
	}

	var k = 100000
	var v = 0
	for key, value := range m {
		// 如果 value 值大于 记录值 v 或者 相等得时候, key更小
		if value > v || value == v && key < k {
			v = value
			k = key
			ans = key
		}
	}
	return
}
