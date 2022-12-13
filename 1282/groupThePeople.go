package groupThePeople

func groupThePeople(groupSizes []int) (ans [][]int) {
	groups := map[int][]int{}

	// 遍历 groupThePeople
	for i, val := range groupSizes {
		groups[val] = append(groups[val], i)
	}

	for size, vals := range groups {
		for index := 0; index < len(vals); index += size {
			ans = append(ans, vals[index: index+size])
		}
	}
	return
}
