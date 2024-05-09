package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := Person{"Alice", 20}
	fmt.Println(p1)

	bytes, _ := json.Marshal(p1)
	fmt.Println(string(bytes))
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
