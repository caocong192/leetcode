package fourSum

import (
	"reflect"
	"testing"
)

//  TestFourSum 单元测试
func TestFourSum(t *testing.T) {
	type test struct { // 定义test结构体
		n      []int
		target int
		want   [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		"exampleTest2": {[]int{1, 2, 3}, 5, [][]int{}},
		"exampleTest3": {[]int{}, 0, [][]int{}},
		"exampleTest4": {[]int{1, -2, -5, -4, -3, 3, 3, 5}, -11, [][]int{{-5, -4, -3, 1}}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := fourSum(ts.n, ts.target)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
