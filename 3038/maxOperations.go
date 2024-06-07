package maxOperations

func maxOperations(nums []int) (ans int) {

	sum := nums[0] + nums[1]
	ans = 1

	for i := 2; i < len(nums)-1; i += 2 {
		if sum == nums[i]+nums[i+1] {
			ans++
			continue
		}
		return
	}
	return
}
