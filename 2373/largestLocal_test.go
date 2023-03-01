package largestLocal

import (
	"reflect"
	"testing"
)

// TestLargestLocal 单元测试
func TestLargestLocal(t *testing.T) {
	type test struct { // 定义test结构体
		grid [][]int
		want [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}, [][]int{{9, 9}, {8, 6}}},
		"exampleTest2": {[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}, [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := largestLocal(ts.grid)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
