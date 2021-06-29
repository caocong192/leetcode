package convertToTitle

func convertToTitle(columnNumber int) (ans string) {
	charSlice := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for columnNumber > 0 {
		if columnNumber%26 == 0 {
			ans = charSlice[25] + ans
			columnNumber -= 26
		} else {
			ans = charSlice[columnNumber%26-1] + ans
		}
		columnNumber /= 26
	}
	return
}
