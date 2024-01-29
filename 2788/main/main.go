package main

import "fmt"

func main() {
	s := "one"
	for i, v := range s {
		fmt.Printf("i:%v, v:%v\n", i, string(v))
	}

}
