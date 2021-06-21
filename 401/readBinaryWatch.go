package readBinaryWatch

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	// 初始化 字符串切片
	ret := make([]string, 0)

	// 不存在大于8个指示灯亮的情况, 直接返回空字符串切片
	if turnedOn > 8 {
		return ret
	}

	// 遍历 0-11 小时的情况
	for h := uint8(0); h < 12; h++ {
		// 遍历 0-59 分钟的情况
		for m := uint8(0); m < 60; m++ {
			// 统计小时和分钟的二进制 1位数个数, 是否等于 turnOn
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ret = append(ret, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ret
}
