package minTaps

import (
	"reflect"
	"testing"
)

// TestMinTaps 单元测试
func TestMinTaps(t *testing.T) {
	type test struct { // 定义test结构体
		n      int
		ranges []int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {5, []int{3, 4, 1, 1, 0, 0}, 1},
		"exampleTest2": {3, []int{0, 0, 0, 0}, -1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := minTaps(ts.n, ts.ranges)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
