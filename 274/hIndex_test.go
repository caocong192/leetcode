package hIndex

import (
	"reflect"
	"testing"
)

// TestHIndex 单元测试
func TestHIndex(t *testing.T) {
	type test struct { // 定义test结构体
		citations []int
		want      int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{3, 0, 6, 1, 5}, 3},
		"exampleTest2": {[]int{1, 3, 1}, 1},
		"exampleTest3": {[]int{0}, 0},
		"exampleTest4": {[]int{10}, 1},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := hIndex(ts.citations)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
