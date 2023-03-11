package findLongestSubarray

func findLongestSubarray(array []string) []string {

	//	使用哈希表记录每个前缀和第一次出现的下标。
	index := make(map[int]int)

	maxLength := 0
	index[0] = -1
	startIndex := -1
	sum := 0
	for i, v := range array {
		if v[0] >= '0' && v[0] <= '9' {
			sum++
		} else {
			sum--
		}

		// 如果 哈希表index 之前有值, 则和 maxLength进行比较
		if firstIndex, ok := index[sum]; ok {
			if i-firstIndex > maxLength {
				maxLength = i - firstIndex
				startIndex = firstIndex + 1
			}
		} else {
			index[sum] = i
		}

		//fmt.Printf("当前前缀和:%v, startIndex:%v, maxLength:%v\n", sum, startIndex, maxLength)
	}
	//
	// 如果最后 sum 值为0, 则最大的字串就是本身
	if sum == 0 {
		return array
	}

	// 如果最大长度还是0, 则没有合适的字串
	if maxLength == 0 {
		return []string{}
	}

	return array[startIndex : startIndex+maxLength]
}
