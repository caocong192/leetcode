package convertToBase7

import "strconv"

//func convertToBase7(num int) string {
//	if num == 0{
//		return "0"
//	}
//
//	ans := []byte{}
//	flag := num < 0
//	if flag {
//		num = -num
//	}
//
//	for num >0 {
//		ans = append(ans, '0'+byte(num%7))
//		num /=7
//	}
//
//	if flag {
//		ans = append(ans, '-')
//	}
//
//	for i, n :=0, len(ans); i<n/2; i++{
//		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
//	}
//	return  string(ans)
//}


func convertToBase7(num int) string {
	if num == 0 {
		return  "0"
	}

	flag := num < 0
	if flag {
		num = -num
	}

	ans := ""
	for num >0 {
		tail := num %7
		num /=7
		ans = strconv.Itoa(tail) + ans
	}

	if flag {
		ans = "-" + ans
	}

	return ans
}