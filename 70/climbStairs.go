package climbStairs

func climbStairs(n int) int {
	//return f(n)

	var m = make(map[int]int)
	m[1] = 1
	m[2] = 2
	for i := 3; i <= 45; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}

//
//func f(i int) (ans int) {
//	if i == 1 {
//		return 1
//	}
//
//	if i == 2 {
//		return 2
//	}
//
//	return f(i-1) + f(i-2)
//}
