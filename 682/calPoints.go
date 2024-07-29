package calPoints

import "strconv"

func calPoints(operations []string) int {
	for i := 0; i < len(operations); i++ {
		if operations[i] == "+" {
			v1, _ := strconv.Atoi(operations[i-2])
			v2, _ := strconv.Atoi(operations[i-1])
			operations[i] = strconv.Itoa(v1 + v2)
		} else if operations[i] == "D" {
			v, _ := strconv.Atoi(operations[i-1])
			operations[i] = strconv.Itoa(v * 2)
		} else if operations[i] == "C" {
			operations = append(operations[:i-1], operations[i+1:]...)
			i -= 2
		}
	}

	var ans int
	for i := 0; i < len(operations); i++ {
		v, _ := strconv.Atoi(operations[i])
		ans += v
	}
	return ans
}
