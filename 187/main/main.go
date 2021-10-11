package main

import "fmt"

func main() {
	s := "AAAAAAAAAA"
	ret := findRepeatedDnaSequences(s)
	fmt.Println(ret)

}

func findRepeatedDnaSequences(s string) (want []string) {
	n := len(s)
	if n < 10 {
		return want
	}

	var hashMap = make(map[string]int)

	for i := 0; i <= n-10; i++ {
		currStr := s[i : i+10]
		if _, ok := hashMap[currStr]; ok {
			hashMap[currStr] ++
		} else {
			hashMap[currStr] = 1
		}
	}

	for k, v := range hashMap {
		if v > 1 {
			want = append(want, k)
		}
	}
	return
}
