package countWords

import "fmt"

func countWords(words1 []string, words2 []string) (res int) {
	wordHash := map[string]int{}

	for _, w := range words1 {
		wordHash[w]++
	}

	for _, w := range words2 {
		if r, ok := wordHash[w]; ok && r == 1 {
			res++
			wordHash[w] = -1
			fmt.Println("res:", res)
		} else if r, ok = wordHash[w]; ok && r == -1 {
			res--
			wordHash[w] = -2
			fmt.Println("res:", res)
		}
	}
	return res
}
