package arrangeCoins

import (
	"sort"
)

//func arrangeCoins(n int) int {
//	for i := math.Sqrt(float64(n)); ; i++ {
//		if int(i)*(int(i)+1) == n*2 {
//			return int(i)
//		} else if int(i)*(int(i)+1) > n*2 {
//			return int(i) - 1
//		}
//	}
//	return 0
//}

func arrangeCoins(n int) int {
	return sort.Search(n, func(i int) bool {
		i++
		return i*(1+i) > 2*n
	})
}
