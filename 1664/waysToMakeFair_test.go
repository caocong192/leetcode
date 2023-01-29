package waysToMakeFair

import (
	"reflect"
	"testing"
)

// TestWaysToMakeFair 单元测试
func TestWaysToMakeFair(t *testing.T) {
	type test struct { // 定义test结构体
		nums []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 1, 6, 4}, 1},
		"exampleTest2": {[]int{1, 1, 1}, 3},
		"exampleTest3": {[]int{1, 2, 3}, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := waysToMakeFair(ts.nums)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
