package printBin

func printBin(num float64) string {

	var a = 0.5
	ans := "0."
	for num != 0 {
		if num >= a {
			num -= a
			ans += "1"
		} else {
			ans += "0"
		}
		a /= 2

		if len(ans) > 8 {
			return "ERROR"
		}
	}
	return ans
}
