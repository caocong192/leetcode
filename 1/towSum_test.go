package twoSum

import (
	"reflect"
	"testing"
)

// TestTwoSum 单元测试
func TestTwoSum(t *testing.T) {
	type test struct { // 定义test结构体
		arr    []int
		target int
		want   []int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		"exampleTest2": {[]int{3,2,4}, 6, []int{1,2}},
		"exampleTest3": {[]int{3,3}, 6, []int{0,1}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := twoSum(ts.arr, ts.target)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
