package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s) // s := "pwwkew"
	sLen := len(runes) // sLen := 6

	for k := sLen; k > 1; k-- {
		for i := 0; i+k <= sLen; i++ {
			j := i + k
			newRunes := runes[i:j]
			for index, val := range newRunes {
				for index2, val2 := range newRunes {
					if val == val2 && index != index2 {
						goto f
					}
				}
				if index == len(newRunes)-1 {
					fmt.Println(i, j)
					return len(newRunes)
				}
			}
		f:
		}
	}
	return 1
}

func main() {
	s := "au"

	ret := lengthOfLongestSubstring(s)
	fmt.Println(ret)

}
