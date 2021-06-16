package stoneGame

import (
	"reflect"
	"testing"
)

// TestPeakIndexInMountain 单元测试
func TestStoneGame(t *testing.T) {
	type test struct { // 定义test结构体
		arr  []int
		want bool
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{5, 3, 4, 5}, true},
		"exampleTest2": {[]int{5, 1, 1, 1, 1, 1, 9, 5}, true},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := stoneGame(ts.arr)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
