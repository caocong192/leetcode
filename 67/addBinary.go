package addBinary

import (
	"strconv"
)

func addBinary(a string, b string) string {
	a, b = reverse(a), reverse(b)
	flag := 0
	c := ""
	for i := 0; i < len(a) || i < len(b); i++ {
		sum := 0
		if i < len(a) {
			//res, _ := strconv.Atoi(string(a[i]))
			//sum += res

			sum += int(a[i] - '0')
		}

		if i < len(b) {
			//res, _ := strconv.Atoi(string(b[i]))
			//sum += res
			sum += int(b[i] - '0')

		}
		sum += flag
		flag = sum / 2
		sum = sum % 2
		c += strconv.Itoa(sum)
	}

	if flag != 0 {
		c += strconv.Itoa(flag)
	}
	return reverse(c)
}

func reverse(a string) (b string) {
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	return
}
