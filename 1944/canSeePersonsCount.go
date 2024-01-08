package canSeePersonsCount

func canSeePersonsCount(heights []int) []int {

	// 单调栈
	n := len(heights)
	stack := make([]int, 0)
	want := make([]int, len(heights))

	for i := n - 1; i >= 0; i-- {
		h := heights[i]
		// 当前元素比栈顶大, 视为找到下一个更大元素, 就弹出栈顶, 更新结果
		for len(stack) > 0 && stack[len(stack)-1] <= h {
			stack = stack[:len(stack)-1]
			// 更新结果
			want[i]++
		}
		if len(stack) > 0 {
			want[i]++
		}
		// 存入栈
		stack = append(stack, h)
	}
	return want
}
