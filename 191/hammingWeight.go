package hammingWeight

func hammingWeight(num uint32) (ans int) {
	for num > 0 {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return
}
