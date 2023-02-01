package checkXMatrix

import (
	"reflect"
	"testing"
)

// TestCheckXMatrix 单元测试
func TestCheckXMatrix(t *testing.T) {
	type test struct { // 定义test结构体
		grid [][]int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}, true},
		"exampleTest2": {[][]int{{5, 7, 0}, {0, 3, 1}, {0, 5, 0}}, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := checkXMatrix(ts.grid)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
