package diagonalSort

import (
	"reflect"
	"testing"
)

// TestDiagonalSort 单元测试
func TestDiagonalSort(t *testing.T) {
	type test struct { // 定义test结构体
		mat  [][]int
		want [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}, [][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}}},
		"exampleTest2": {[][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}},
			[][]int{{5, 17, 4, 1, 52, 7}, {11, 11, 25, 45, 8, 69}, {14, 23, 25, 44, 58, 15}, {22, 27, 31, 36, 50, 66}, {84, 28, 75, 33, 55, 68}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := diagonalSort(ts.mat)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
