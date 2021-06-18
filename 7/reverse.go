package reverse

func reverse(x int) int {
	// 定义一个整数 ret 用于存储反转后的值
	ret := 0

	for {
		// 一直遍历到x为0
		if x != 0 {
			// 一直判断到倒数第二次计算时候, ret结果是否越界.   因为x 为int类型, 第一位数<=2, 临界值最后一位不大于7.  所以最后一次不用判断
			if ret < (-1<<31)/10 || ret > (1<<31-1)/10 {
				return 0
			}
			digit := x % 10
			x = x / 10
			ret = ret*10 + digit
		} else {
			return ret
		}
	}
}
