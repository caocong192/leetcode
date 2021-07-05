package maxPoints

import (
	"reflect"
	"testing"
)

//  TestMaxPoints 单元测试
func TestMaxPoints(t *testing.T) {
	type test struct { // 定义test结构体
		points [][]int
		want   int
	}
	tests := map[string]test{ // 测试用例使用map存储
		"exampleTest1": {[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		"exampleTest2": {[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		"exampleTest3": {[][]int{{0, 1}, {0, 0}}, 2},
		"exampleTest4": {[][]int{{0, 0}}, 1},
		"exampleTest5": {[][]int{{4, 5}, {4, 1}, {4, 2}, {4, 4}}, 4},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := maxPoints(ts.points)
			if !reflect.DeepEqual(got, ts.want) {
				t.Errorf("excepted:%#v, got:%#v", ts.want, got)
			}
		})
	}
}
