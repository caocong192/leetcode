package numberOfRightTriangles

import (
	"reflect"
	"testing"
)

// TestNumberOfRightTriangles
func TestNumberOfRightTriangles(t *testing.T) {
	type test struct { // 定义test结构体
		grid [][]int
		want int64
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}, 2},
		"exampleTest2": {[][]int{{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0}}, 0},
		"exampleTest3": {[][]int{{0, 0, 0}}, 0},
		"exampleTest4": {[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 0}}, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := numberOfRightTriangles(ts.grid)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
