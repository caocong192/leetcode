package minRectanglesToCoverPoints

import (
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	start := points[0][0]
	ans++
	for i := 1; i < len(points); i++ {
		if points[i][0]-start <= w {
			continue
		}
		ans++
		start = points[i][0]
	}
	return
}
