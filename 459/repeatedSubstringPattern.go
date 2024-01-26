package repeatedSubstringPattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		repeated := strings.Repeat(s[:i], n/i)
		if repeated == s {
			return true
		}
	}
	return false
}
