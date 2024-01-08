package numOfBurgers

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	//巨无霸汉堡：4 片番茄和 1 片奶酪
	//小皇堡：2 片番茄和 1 片奶酪

	// 如果 番茄切片是单数、 全做小皇堡 番茄还少了、 全做巨无霸汉堡 番茄还多了, 则不可能不剩原料
	if tomatoSlices%2 == 1 || tomatoSlices < cheeseSlices*2 || tomatoSlices > cheeseSlices*4 {
		return []int{}
	}

	x := tomatoSlices/2 - cheeseSlices
	y := cheeseSlices - x
	return []int{x, y}
}
