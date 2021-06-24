package hammingWeight

import "math/bits"

func hammingWeight(num uint32) int {
	ans := bits.OnesCount32(num)
	return ans
}
