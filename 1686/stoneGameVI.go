package stoneGameVI

import (
	"fmt"
	"sort"
)

func stoneGameVI(aliceValues []int, bobValues []int) int {

	n := len(aliceValues)
	values := make([][]int, n)
	for i := 0; i < n; i++ {
		values[i] = []int{aliceValues[i] + bobValues[i], aliceValues[i], bobValues[i]}
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i][0] > values[j][0]
	})

	res := 0
	fmt.Println(values)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			res += values[i][1]
		} else {
			res -= values[i][2]
		}
	}

	if res > 0 {
		return 1
	} else if res == 0 {
		return 0
	} else {
		return -1
	}
}
