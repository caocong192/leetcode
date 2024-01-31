package minimumSeconds

func minimumSeconds(nums []int) (ans int) {
	mp := make(map[int][]int)
	n := len(nums)
	ans = n

	for i, num := range nums {
		mp[num] = append(mp[num], i)
	}

	for _, pos := range mp {
		mx := pos[0] + n - pos[len(pos)-1]
		for i := 1; i < len(pos); i++ {
			mx = max(mx, pos[i]-pos[i-1])
		}
		ans = min(ans, mx/2)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
