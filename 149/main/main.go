package main

import "fmt"

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	var para [][]float32
	var m = map[int]int{}

	for index1, point1 := range points {
		m[point1[0]] += 1

		for index2, point2 := range points {
			if index1 != index2 && point1[0] != point2[0] {
				a := float32(point2[1]-point1[1]) / float32(point1[0]-point2[0])
				b := float32(point1[0])*a + float32(point1[1])
				para = append(para, []float32{a, b})
			}
		}
	}

	ans := 0
	for _, p := range para {
		maxPoint := 0
		for _, point := range points {
			if float32(point[0])*p[0]+float32(point[1]) == p[1] {
				maxPoint++
			}
		}

		if ans < maxPoint {
			ans = maxPoint
		}
	}

	for _, point := range points {
		if m[point[0]] > ans {
			ans = m[point[0]]
		}
	}

	return ans
}

func main() {
	points := [][]int{{-6, -1}, {3, 1}, {12, 3}}
	ret := maxPoints(points)
	fmt.Println(ret)

}
