package minimumRounds

func minimumRounds(tasks []int) (ans int) {
	// hash 表
	hash := make(map[int]int)
	for _, v := range tasks {
		hash[v]++
	}

	for _, v := range hash {
		if v == 1 {
			return -1
		}
	}

	for _, v := range hash {
		ans += getGround(v)
	}
	return
}

func getGround(count int) int {
	if count%3 == 0 {
		return count / 3
	}
	return count/3 + 1
}
