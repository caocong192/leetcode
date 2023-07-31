package maxSumDivThree

func maxSumDivThree(nums []int) (ans int) {
	remainder := make(map[int]int)
	for _, v := range nums {
		a := remainder[0] + v
		b := remainder[1] + v
		c := remainder[2] + v

		remainder[a%3] = max(remainder[a%3], a)
		remainder[b%3] = max(remainder[b%3], b)
		remainder[c%3] = max(remainder[c%3], c)
	}

	return remainder[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
