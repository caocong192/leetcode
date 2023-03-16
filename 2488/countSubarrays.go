package countSubarrays

func countSubarrays(nums []int, k int) int {

	// 前缀和 + 哈希表
	var m = make(map[int]int)

	var kIndex int
	for i, v := range nums {
		if v == k {
			kIndex = i
		}
	}

	var sum, ans int
	m[0] = 1
	for i, v := range nums {
		sum += getValue(v - k)
		if i < kIndex {
			m[sum]++
		} else {
			ans += m[sum] + m[sum-1]
		}

	}
	return ans
}

func getValue(n int) int {
	if n > 0 {
		return 1
	} else if n == 0 {
		return 0
	}
	return -1
}
