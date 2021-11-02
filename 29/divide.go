package divide

import "math"

func divide(dividend int, divisor int) (ans int) {
	// 如果被除数为 math.MinInt32 并且除数为  -1
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// 如果被除数为0
	if dividend == 0 {
		return 0
	}

	// 获取最终符号位, 并将被除数和除数取正
	symbol := 1
	if dividend < 0 {
		symbol *= -1
		dividend *= -1
	}
	if divisor < 0 {
		symbol *= -1
		divisor *= -1
	}

	// 位运算
	for i := 32; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			ans += 1 << i
			dividend -= divisor << i
		}
	}

	ans *= symbol
	return
}
