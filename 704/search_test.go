package search

import (
	"reflect"
	"testing"
)

// TestSearch 单元测试
func TestSearch(t *testing.T) {
	type test struct { // 定义test结构体
		nums   []int
		target int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		"exampleTest2": {[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		"exampleTest3": {[]int{-5, 0, 3, 5, 9, 12, 18}, -5, 0},
		"exampleTest4": {[]int{-1, 0, 3, 5, 9, 12, 19, 21}, 21, 7},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := search(ts.nums, ts.target)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
