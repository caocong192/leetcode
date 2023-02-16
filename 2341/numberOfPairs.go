package numberOfPairs

func numberOfPairs(nums []int) []int {

	// 使用map 存储对应的 数值 和 个数
	var m = map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}

	// 遍历map, 获取每个 数值 的配对数 和 遗留数
	pair, left := 0, 0
	for _, v := range m {
		pair += v / 2
		left += v % 2
	}

	return []int{pair, left}
}
