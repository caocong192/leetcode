package searchInsert

import (
	"reflect"
	"testing"
)

// TestSearchInsert 单元测试
func TestSearchInsert(t *testing.T) {
	type test struct { // 定义test结构体
		nums   []int
		target int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 3, 5, 6}, 5, 2},
		"exampleTest2": {[]int{1, 3, 5, 6}, 2, 1},
		"exampleTest3": {[]int{1, 3, 5, 6}, 7, 4},
		"exampleTest4": {[]int{2, 3, 5, 6}, 1, 0},
	}
	for name, ts := range tests {
		t.Run(name, func(s *testing.T) { // 使用t.Run()执行子测试
			got := searchInsert(ts.nums, ts.target)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
