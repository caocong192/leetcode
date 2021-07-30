package titleToNumber

func titleToNumber2(columnTitle string) (ans int) {
	for i, mul := len(columnTitle)-1, 1; i >= 0; i-- {
		num := columnTitle[i] - 'A' + 1
		ans = int(num)*mul + ans
		mul *= 26
	}
	return
}
