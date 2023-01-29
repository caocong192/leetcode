package validPath

import (
	"reflect"
	"testing"
)

// TestValidPath 单元测试
func TestValidPath(t *testing.T) {
	type test struct { // 定义test结构体
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2, true},
		"exampleTest2": {6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5, false},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := validPath(ts.n, ts.edges, ts.source, ts.destination)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
