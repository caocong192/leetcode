package maxArea

func maxArea(height []int) int {
	// 定义最大的 area
	area := 0

	// left, right 分别表示数组左右下标
	left, right := 0, len(height)-1

	// h 定义计算面积的时候的高度
	h := 0

	// 当左小标 小于 右下标 一直循环
	for left < right {
		// t 为当前的长度
		t := right - left
		// h 高度取当前两个值的最小值, 如果left小, 则left向右移动, 如果right小, 则right向左移动
		if height[left] <= height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		// area 保存最大的面积
		if area < h*t {
			area = h * t
		}
	}
	return area
}
