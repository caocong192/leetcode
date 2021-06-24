package maxArea

import (
	"reflect"
	"testing"
)

//  TestMaxArea 单元测试
func TestMaxArea(t *testing.T) {
	type test struct { // 定义test结构体
		s    []int
		want int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		"exampleTest2": {[]int{1, 1}, 1},
		"exampleTest3": {[]int{4, 3, 2, 1, 4}, 16},
		"exampleTest4": {[]int{1, 2, 1}, 2},
		"exampleTest5": {[]int{1, 2, 3, 4, 5, 21, 21, 5, 4, 3, 2, 1}, 21},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maxArea(ts.s)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
