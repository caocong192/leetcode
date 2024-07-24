package countBeautifulPairs

import (
	"fmt"
	"strconv"
)

func countBeautifulPairs(nums []int) (ans int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			x, _ := strconv.Atoi(fmt.Sprintf("%d", nums[i])[:1])
			y := nums[j] % 10

			if x < y {
				x, y = y, x
			}
			fmt.Printf("x:%d,y:%d\n", x, y)
			if gcdCircle(x, y) == 1 {
				ans++
			}
		}
	}
	return
}

func gcdCircle(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		gcdCircle(y, tmp)
	} else {
		return y
	}
	return 0
}
