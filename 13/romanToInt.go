package romanToInt

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(s string) (ans int) {

	for len(s) > 0 {
		// 标识两个罗马字符的情况是否匹配到了
		f := true

		// 判断二个罗马数字情况
		if len(s) >= 2 {
			for _, val := range valueSymbols {
				if s[0:2] == val.symbol {
					ans += val.value
					if len(s) == 2 {
						return
					}
					f = false
					s = s[2:]
					f = false
					break
				}
			}
		}

		// 判断一个罗马数字情况
		if f {
			for _, val := range valueSymbols {
				if s[0:1] == val.symbol {
					ans += val.value
					if len(s) == 1 {
						return
					}
					s = s[1:]
					break
				}
			}
		}

	}
	return
}
