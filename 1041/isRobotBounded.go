package isRobotBounded

func isRobotBounded(instructions string) bool {
	var direction = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}

	var step = 0
	var dire = 0

	// 循环4次, 即可一个轮回判断是否满足
	for i := 1; i <= 4; i++ {
		for _, s := range instructions {

			if s == 'G' {
				step++
				continue
			}

			if s == 'L' {
				direction[dire] += step
				dire = (dire + 3) % 4
			} else {
				direction[dire] += step
				dire = (dire + 1) % 4
			}
			step = 0
		}
	}

	direction[dire] += step

	return direction[0] == direction[2] && direction[1] == direction[3]
}
