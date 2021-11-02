package divide

import "math"

func divide2(dividend int, divisor int) (ans int) {
	// 如果被除数为 math.MinInt32 并且除数为  -1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// 如果被除数为0
	if dividend == 0 {
		return 0
	}

	// 获取最终符号位, 并将被除数和除数取正
	symbol := (dividend ^ divisor) < 0
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	// 位运算
	for i := 32; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			ans += 1 << i
			dividend -= divisor << i
		}
	}

	if symbol {
		return ans * -1
	}
	return
}
