package pathInZigZagTree

func pathInZigZagTree(label int) []int {

	// 计算行号
	row, num := 1, 1
	for num*2 <= label {
		row++
		num *= 2
	}

	var path []int

	// 加入label本身
	path = append(path, label)
	label = getRevLabel(row, label)
	row--

	for label > 1 {
		parentLabel := getRevLabel(row, label/2)
		path = append(path, parentLabel)
		label >>=1
		row--
	}
	return rev(path)
}

// 如果该行是 偶数行, 则反转label
func getRevLabel(row, label int) int {
	if row%2 == 0 {
		return 1<<(row-1) + 1<<row - 1 - label
	}
	return label
}

// 逆序切片
func rev(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
