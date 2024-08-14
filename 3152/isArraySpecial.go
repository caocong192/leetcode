package isArraySpecial

func isArraySpecial(nums []int, queries [][]int) (want []bool) {
	// 使用动态规划, 统计每个数字作为最后一个, 前面有多少符合要求的
	var dp = make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		if (nums[i]+nums[i-1])%2 == 1 {
			dp[i] += dp[i-1]
		} else {
			dp[i] = 0
		}
	}

	for i := 0; i < len(queries); i++ {
		want = append(want, dp[queries[i][1]] >= queries[i][1]-queries[i][0])
	}

	return
}

//func isArraySpecial(nums []int, queries [][]int) (want []bool) {
//	var li [][]int
//	for i := 0; i < len(nums)-1; i++ {
//		j := i + 1
//		for ; j < len(nums); j++ {
//			if (nums[j]+nums[j-1])%2 == 0 {
//				break
//			}
//		}
//		li = append(li, []int{i, j - 1})
//		i = j - 1
//	}
//
//	//fmt.Println(li)
//	for _, v := range queries {
//		l := len(want)
//		if v[0] == v[1] {
//			want = append(want, true)
//		} else {
//			for _, vv := range li {
//				if v[0] > vv[1] {
//					continue
//				}
//				// 如果vv[0]>v[1] 说明vv[0]在v[1]的右边, 则后面的都不需要比对
//				if v[1] < vv[0] {
//					break
//				}
//				if vv[0] <= v[0] && vv[1] >= v[1] {
//					want = append(want, true)
//					break
//				}
//
//			}
//			if l == len(want) {
//				want = append(want, false)
//			}
//		}
//	}
//	return
//}
