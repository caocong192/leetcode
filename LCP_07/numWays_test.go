package numWays

import (
	"reflect"
	"testing"
)

// TestNumWays 单元测试
func TestNumWays(t *testing.T) {
	type test struct { // 定义test结构体
		n        int
		relation [][]int
		k        int
		want     int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {5, [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 3, 3},
		"exampleTest2": {3, [][]int{{0, 2}, {2, 1}}, 2, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := numWays(ts.n, ts.relation, ts.k)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
