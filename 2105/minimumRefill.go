package minimumRefill

func minimumRefill(plants []int, capacityA int, capacityB int) (ans int) {
	// 双指针
	for i, j, a, b := 0, len(plants)-1, capacityA, capacityB; i <= j; {
		if i == j {
			if a > b {
				if a < plants[i] {
					ans++
				}
			} else {
				if b < plants[i] {
					ans++
				}
			}
			break
		}

		if a < plants[i] {
			a = capacityA
			ans++
		}
		if b < plants[j] {
			b = capacityB
			ans++
		}
		a -= plants[i]
		b -= plants[j]
		i++
		j--
	}
	return
}
