package hardestWorker

import (
	"reflect"
	"testing"
)

// TestHardestWorker 单元测试
func TestHardestWorker(t *testing.T) {
	type test struct { // 定义test结构体
		n    int
		logs [][]int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}, 3},
		"exampleTest2": {2, [][]int{{0, 10}, {1, 20}}, 0},
		"exampleTest3": {10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := hardestWorker(ts.n, ts.logs)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
