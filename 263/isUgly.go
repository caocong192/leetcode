package isUgly

func isUgly(n int) bool {

	if n < 1 {
		return false
	}

	nums := []int{2, 3, 5}

	for _, f := range nums {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}
