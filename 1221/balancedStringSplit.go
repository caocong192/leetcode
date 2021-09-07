package balancedStringSplit

func balancedStringSplit(s string) int {

	strValue := 0
	nums := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "L" {
			strValue--
		} else {
			strValue++
		}

		if strValue == 0 {
			nums++
		}
	}
	return nums
}
