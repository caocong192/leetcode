package peakIndexInMountain

import (
	"reflect"
	"testing"
)

// TestPeakIndexInMountain 单元测试
func TestPeakIndexInMountain(t *testing.T) {
	type test struct { // 定义test结构体
		arr  []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{0, 1, 0}, 1},
		"exampleTest2": {[]int{0, 2, 1, 0}, 1},
		"exampleTest3": {[]int{0, 10, 5, 2}, 1},
		"exampleTest4": {[]int{3, 4, 5, 1}, 2},
		"exampleTest5": {[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
		"exampleTest6": {[]int{0, 10, 11, 12, 13, 14, 15, 2}, 6},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := peakIndexInMountainArray(ts.arr)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
