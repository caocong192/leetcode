package restoreMatrix

import (
	"reflect"
	"testing"
)

// TestRestoreMatrix 单元测试
func TestRestoreMatrix(t *testing.T) {
	type test struct { // 定义test结构体
		rowSum []int
		colSum []int
		want   [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3, 8}, []int{4, 7}, [][]int{{3, 0}, {1, 7}}},
		"exampleTest2": {[]int{5, 7, 10}, []int{8, 6, 8}, [][]int{{0, 5, 0}, {6, 1, 0}, {2, 0, 8}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := restoreMatrix(ts.rowSum, ts.colSum)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
