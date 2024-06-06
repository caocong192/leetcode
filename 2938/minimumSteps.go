package minimumSteps

func minimumSteps(s string) (ans int64) {
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			sum++
		} else {
			ans += int64(sum)
		}
	}
	return
}
