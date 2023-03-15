package generate

import (
	"reflect"
	"testing"
)

// TestGenerate
func TestGenerate(t *testing.T) {
	type test struct { // 定义test结构体
		numRows int
		want    [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		"exampleTest2": {1, [][]int{{1}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := generate(ts.numRows)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
