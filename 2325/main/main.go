package main

import "fmt"

func main() {

	key := "the quick brown fox jumps over the lazy dog"
	m := make(map[string]string)
	x := 97
	res := make(map[string]string)

	for _, v := range key {
		if string(v) == " " || m[string(v)] != "" {
			continue
		}
		m[string(v)] = string(v)
		res[string(x)] = string(v)
		x++

	}

	for k, v := range res {
		fmt.Println(k, v)
	}

}
