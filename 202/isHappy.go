package isHappy

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)

	for {
		n = multiplyNum(n)
		fmt.Println("res:", n)
		if n == 1 {
			return true
		}

		if m[n] == true {
			return false
		} else {
			m[n] = true
		}
	}
}

func multiplyNum(num int) (ans int) {
	for num > 0 {
		n := num % 10
		ans += n * n
		num /= 10
	}
	return
}
