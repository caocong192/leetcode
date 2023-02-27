package mySqrt

func mySqrt(x int) int {

	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans

	//for i := 0; i < math.MaxInt32; i++ {
	//	if i*i == x {
	//		return i
	//	}
	//
	//	if i*i > x {
	//		return i - 1
	//	}
	//}
	//return 0
}
