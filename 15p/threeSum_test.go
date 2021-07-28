package threeSum

import (
	"reflect"
	"testing"
)

//  TestThreeSum 单元测试
func TestThreeSum(t *testing.T) {
	type test struct { // 定义test结构体
		n    []int
		want [][]int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		"exampleTest2": {[]int{0, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		"exampleTest3": {[]int{}, [][]int{}},
		"exampleTest4": {[]int{0, 1}, [][]int{}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := threeSum(ts.n)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
