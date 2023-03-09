package minimumRecolors

func minimumRecolors(blocks string, k int) (ans int) {
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			ans++
		}
	}

	if ans == 0 || len(blocks) == k {
		return
	}

	res := ans

	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			res++
		}

		if blocks[i-k] == 'W' {
			res--
		}

		if res < ans {
			ans = res
		}
	}
	return
}
