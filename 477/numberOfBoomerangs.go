package numberOfBoomerangs

func numberOfBoomerangs(points [][]int) (res int) {

	// 少于3个点, 无法构成回旋镖最小数量, 直接返回0
	if len(points) < 3 {
		return 0
	}

	// 遍历每个点 和其它点的距离
	for i := 0; i < len(points); i++ {
		// hash表 记录所有点于其它点的距离
		distances := map[int]int{}
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			d := (points[j][0]-points[i][0])*(points[j][0]-points[i][0]) + (points[j][1]-points[i][1])*(points[j][1]-points[i][1])
			distances[d]++
		}

		// 结果大于2, 那么以该点为顶点, 可以组成 n*(n-1) 个回旋镖
		for _, v := range distances {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}
	return
}
