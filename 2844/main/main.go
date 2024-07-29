package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 2, 2, 1, 4, 5, 3, 4}
	sort.Ints(nums)
	fmt.Println(nums)

}
