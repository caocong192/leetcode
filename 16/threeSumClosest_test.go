package threeSumClosest

import (
	"reflect"
	"testing"
)

//  TestThreeSumClosest 单元测试
func TestThreeSumClosest(t *testing.T) {
	type test struct { // 定义test结构体
		n      []int
		target int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[]int{-1, 2, 1, -4}, 1, 2},
		"exampleTest2": {[]int{0, 0, 0, 0}, 0, 0},
		"exampleTest3": {[]int{1, 2, 3, 4, 5}, 12, 12},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := threeSumClosest(ts.n, ts.target)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
