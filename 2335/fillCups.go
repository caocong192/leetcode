package fillCups

func fillCups(amount []int) (ans int) {

	a := amount[0]
	b := amount[1]
	c := amount[2]

	for ans = 0; a != 0 || b != 0 || c != 0; ans++ {
		if a >= c && b >= c {
			reduce1(&a)
			reduce1(&b)
		} else if a >= b && c >= b {
			reduce1(&a)
			reduce1(&c)
		} else if b >= a && c >= a {
			reduce1(&b)
			reduce1(&c)
		}
	}
	return
}

func reduce1(d *int) {
	if *d > 0 {
		*d -= 1
	}
}
