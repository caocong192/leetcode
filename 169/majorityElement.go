package majorityElement

func majorityElement(nums []int) int {
	n := len(nums)/2 + 1
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		if m[v] >= n {
			return v
		}
	}
	return -1
}
