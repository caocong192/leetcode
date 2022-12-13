package pondSizes

import (
	"reflect"
	"testing"
)

// TestPondSizes 单元测试
func TestPondSizes(t *testing.T) {
	type test struct { // 定义test结构体
		land [][]int
		want []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{0, 2, 1, 0}, {0, 1, 0, 1}, {1, 1, 0, 1}, {0, 1, 0, 1}}, []int{1, 2, 4}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := pondSizes(ts.land)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
