package distinctDifferenceArray

//func distinctDifferenceArray(nums []int) (ans []int) {
//
//	// 哈希表, 存入前缀的元素的下标, 和后缀元素的下标
//	preMp := make(map[int][]int)
//	suffMp := make(map[int][]int)
//	n := len(nums)
//
//	for i := 0; i < n; i++ {
//		suffMp[nums[i]] = append(suffMp[nums[i]], i)
//	}
//
//	for i := 0; i < n; i++ {
//		preMp[nums[i]] = append(preMp[nums[i]], i)
//		if len(suffMp[nums[i]]) == 1 {
//			delete(suffMp, nums[i])
//		} else {
//			suffMp[nums[i]] = suffMp[nums[i]][1:]
//		}
//		ans = append(ans, len(preMp)-len(suffMp))
//	}
//	return
//}

func distinctDifferenceArray(nums []int) (ans []int) {
	// 哈希表 通过结构体存在表示元素存在
	st := map[int]struct{}{}
	n := len(nums)
	sufCnt := make([]int, n+1)

	// 倒序存入 i- n-1 不同元素的个数
	for i := n - 1; i > 0; i-- {
		st[nums[i]] = struct{}{}
		sufCnt[i] = len(st)
	}

	st = map[int]struct{}{}
	for i := 0; i < n; i++ {
		st[nums[i]] = struct{}{}
		ans = append(ans, len(st)-sufCnt[i+1])
	}
	return
}
