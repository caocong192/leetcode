package smallestGoodBase

import (
	"math"
	"math/bits"
	"strconv"
)

// 官方示例
func smallestGoodBase(n string) string {
	// 转成整数类型
	nVal, _ := strconv.Atoi(n)
	// 取出二进制数的长度
	mMax := bits.Len(uint(nVal)) - 1
	for m := mMax; m > 1; m-- {
		// k = n 开根号 m
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == nVal {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(nVal - 1)
}
