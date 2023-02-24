package lengthOfLastWord

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " && ans == 0 {
			continue
		}

		if string(s[i]) == " " {
			break
		}
		ans++

	}
	return ans
}
