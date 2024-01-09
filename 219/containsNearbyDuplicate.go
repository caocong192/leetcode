package containsNearbyDuplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := map[int]int{}

	for i, v := range nums {
		if p, ok := hash[v]; ok && i-p <= k {
			return true
		} else {
			hash[v] = i
		}
	}
	return false
}
